package back_tracking

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
)

type Ticket struct {
	Id       int     //编号
	TicketNo string  //发票号码
	Amount   float64 //金额
}

var IfCalcResult bool

func ParseXlsx(filename string) []Ticket {
	var ticketList []Ticket
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Printf("打开文件失败: %s\n", err)
		return ticketList
	}
	ticketList = []Ticket{}
	// 获取标签页(时间)
	for _, row := range xlFile.Sheets[0].Rows[1:] {
		var strs []string
		t := Ticket{}
		for _, cell := range row.Cells {
			text := cell.String()
			text = strings.Replace(text, " ", "", -1)
			strs = append(strs, text)
		}
		if strs[1] == "" || strs[2] == "" {
			break
		}
		// 获取标签页(时间)
		t.Id, _ = strconv.Atoi(strs[0])
		t.TicketNo = strs[1]
		t.Amount, _ = strconv.ParseFloat(strs[2], 64)
		ticketList = append(ticketList, t)
	}
	return ticketList
}

func CalcTicket(tickets, tmplist []Ticket, start int, target, diff float64) {
	if target < 0 {
		return
	}
	if target <= diff {
		fmt.Println("\n****得到的发票结果为****")
		sum := float64(0)
		for _, t := range tmplist {
			fmt.Printf("%v %v %v\n", t.Id, t.TicketNo, t.Amount)
			sum += t.Amount
		}
		fmt.Printf("\n*****总金额:%v********", sum)
		IfCalcResult = true
		return
	}
	for i := start; i < len(tickets); i++ {
		tmplist = append(tmplist, tickets[i])
		CalcTicket(tickets, tmplist, i+1, target-tickets[i].Amount, diff)
		if IfCalcResult {
			return
		}
		tmplist = tmplist[:len(tmplist)-1]
	}
	return
}
