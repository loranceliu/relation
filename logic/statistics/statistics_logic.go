package statistics

import (
	"fmt"
	"gin/entity/vo"
	"gin/model"
	"gin/svc"
	"strings"
	"time"
)

func GetRelationTypeNum(ctx *svc.ServiceContext) (*[]vo.ChartCommonVO, error) {
	// 构建原生 SQL 查询语句
	sqlQuery := "select count(*) value ,trt.relation_type_name name from tb_relation tr left join tb_relation_type trt on tr.relation_type_id = trt.relation_type_id where tr.owner_id = ? GROUP BY tr.relation_type_id"

	userId := ctx.Value("user_id")

	var pies []vo.ChartCommonVO
	if err := model.DB.Raw(sqlQuery, userId).Scan(&pies).Error; err != nil {
		return nil, err
	}

	return &pies, nil
}

func GetRelationTypeMoney(ctx *svc.ServiceContext) (*[]vo.ChartCommonVO, error) {
	// 构建原生 SQL 查询语句
	sqlQuery := "select sum(tr.money) value ,trt.relation_type_name name from tb_relation tr left join tb_relation_type trt on tr.relation_type_id = trt.relation_type_id where tr.owner_id = ? GROUP BY tr.relation_type_id"

	userId := ctx.Value("user_id")

	var pies []vo.ChartCommonVO
	if err := model.DB.Raw(sqlQuery, userId).Scan(&pies).Error; err != nil {
		return nil, err
	}

	return &pies, nil
}

func GetRelationTotalMoney(ctx *svc.ServiceContext) (*[]vo.ChartCommonVO, error) {
	// 构建原生 SQL 查询语句
	sqlQuery := "select if(transaction_type = 1,'入账','出账') name ,sum(money) value from tb_relation where owner_id = ? GROUP BY transaction_type"

	userId := ctx.Value("user_id")

	var pies []vo.ChartCommonVO

	if err := model.DB.Raw(sqlQuery, userId).Scan(&pies).Error; err != nil {
		return nil, err
	}

	return &pies, nil
}

func GetRelationCurrentYearTrend(ctx *svc.ServiceContext) (*vo.LineChartDataVO, error) {
	// 构建原生 SQL 查询语句
	sqlQuery := "SELECT SUM(money) AS value,if(transaction_type=1,'入账','出账') name,DATE_FORMAT(date, '%c月') AS time FROM tb_relation WHERE owner_id = ? AND YEAR(date) = YEAR(CURDATE()) GROUP BY transaction_type, time"
	userId := ctx.Value("user_id")
	var lines []vo.LineChartCommonVO
	if err := model.DB.Raw(sqlQuery, userId).Scan(&lines).Error; err != nil {
		return nil, err
	}

	var months []string
	for i := 1; i <= 12; i++ {
		months = append(months, fmt.Sprintf("%d月", i))
	}

	lineChart := vo.LineChartDataVO{
		TypeData: []string{"入账", "出账"},
		CateData: months,
	}

	incomeData := vo.SeriesLineChart{
		Name:  "入账",
		Type:  "line",
		Stack: "Total",
		Data:  make([]string, 12),
	}

	expendData := vo.SeriesLineChart{
		Name:  "出账",
		Type:  "line",
		Stack: "Total",
		Data:  make([]string, 12),
	}

	for i := range incomeData.Data {
		incomeData.Data[i] = "0"
	}

	for i := range expendData.Data {
		expendData.Data[i] = "0"
	}

	for ci, cv := range lineChart.CateData {
		for _, lv := range lines {
			if cv == lv.Time && lv.Name == incomeData.Name {
				incomeData.Data[ci] = lv.Value
				continue
			}
		}
	}

	for ci, cv := range lineChart.CateData {
		for _, lv := range lines {
			if cv == lv.Time && lv.Name == expendData.Name {
				expendData.Data[ci] = lv.Value
				continue
			}
		}
	}

	lineChart.SeriesData = append(lineChart.SeriesData, incomeData)
	lineChart.SeriesData = append(lineChart.SeriesData, expendData)

	return &lineChart, nil
}

func GetRelationTenYearTrend(ctx *svc.ServiceContext) (*vo.LineChartDataVO, error) {
	// 构建原生 SQL 查询语句
	sqlQuery := "SELECT SUM(money) AS value,IF(transaction_type = 1, '入账', '出账') AS name,DATE_FORMAT(date, '%Y') AS time FROM tb_relation WHERE owner_id = ? AND YEAR(date) BETWEEN YEAR(CURDATE()) - 9 AND YEAR(CURDATE()) GROUP BY transaction_type, time"
	userId := ctx.Value("user_id")
	var lines []vo.LineChartCommonVO
	if err := model.DB.Raw(sqlQuery, userId).Scan(&lines).Error; err != nil {
		return nil, err
	}

	currentYear := time.Now().Year()
	var recentYears []string
	for i := 0; i < 10; i++ {
		year := currentYear - i
		yearStr := fmt.Sprintf("%d", year)
		recentYears = append(recentYears, yearStr)
	}

	lineChart := vo.LineChartDataVO{
		TypeData: []string{"入账", "出账"},
		CateData: recentYears,
	}

	incomeData := vo.SeriesLineChart{
		Name:  "入账",
		Type:  "line",
		Stack: "Total",
		Data:  make([]string, 10),
	}

	expendData := vo.SeriesLineChart{
		Name:  "出账",
		Type:  "line",
		Stack: "Total",
		Data:  make([]string, 10),
	}

	for i := range incomeData.Data {
		incomeData.Data[i] = "0"
	}

	for i := range expendData.Data {
		expendData.Data[i] = "0"
	}

	for ci, cv := range lineChart.CateData {
		for _, lv := range lines {
			if cv == lv.Time && lv.Name == incomeData.Name {
				incomeData.Data[ci] = lv.Value
				continue
			}
		}
	}

	for ci, cv := range lineChart.CateData {
		for _, lv := range lines {
			if cv == lv.Time && lv.Name == expendData.Name {
				expendData.Data[ci] = lv.Value
				continue
			}
		}
	}

	lineChart.SeriesData = append(lineChart.SeriesData, incomeData)
	lineChart.SeriesData = append(lineChart.SeriesData, expendData)

	return &lineChart, nil
}

func GetRelationUserTopProfit(ctx *svc.ServiceContext, t string) (*vo.ChartBarVO, error) {
	// 构建原生 SQL 查询语句
	var builder strings.Builder
	userId := ctx.Value("user_id")
	if t == "1" {
		_, _ = fmt.Fprintf(&builder, "SELECT tru.relation_user_name AS name,COALESCE(i.total, 0) - COALESCE(e.total, 0) AS value FROM tb_relation_user tru LEFT JOIN ( SELECT SUM(money) AS total, relation_user_id  FROM tb_relation  WHERE transaction_type = 1 AND owner_id = %s GROUP BY relation_user_id) i ON tru.relation_user_id = i.relation_user_id LEFT JOIN ( SELECT SUM(money) AS total, relation_user_id FROM tb_relation WHERE transaction_type = 2 AND owner_id = %s GROUP BY relation_user_id) e ON tru.relation_user_id = e.relation_user_id WHERE owner_id = %s AND (COALESCE(i.total, 0) - COALESCE(e.total, 0)) > 0 ORDER BY value DESC LIMIT 10;", userId, userId, userId)
	}
	if t == "2" {
		_, _ = fmt.Fprintf(&builder, "SELECT tru.relation_user_name AS name,COALESCE(e.total, 0) - COALESCE(i.total, 0) AS value FROM tb_relation_user tru LEFT JOIN ( SELECT SUM(money) AS total, relation_user_id  FROM tb_relation  WHERE transaction_type = 1 AND owner_id = %s GROUP BY relation_user_id) i ON tru.relation_user_id = i.relation_user_id LEFT JOIN ( SELECT SUM(money) AS total, relation_user_id FROM tb_relation WHERE transaction_type = 2 AND owner_id = %s GROUP BY relation_user_id) e ON tru.relation_user_id = e.relation_user_id WHERE owner_id = %s AND (COALESCE(i.total, 0) - COALESCE(e.total, 0)) < 0 ORDER BY value DESC LIMIT 10;", userId, userId, userId)
	}

	sqlQuery := builder.String()
	var pies []vo.ChartCommonVO
	if err := model.DB.Raw(sqlQuery).Scan(&pies).Error; err != nil {
		return nil, err
	}

	bar := &vo.ChartBarVO{}
	var name []string
	var data []string

	if len(pies) > 0 {
		for _, v := range pies {
			name = append(name, v.Name)
			data = append(data, v.Value)
		}
		bar.Name = name
		bar.Data = data
	}

	return bar, nil
}
