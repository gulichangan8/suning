package dao

import (
	"database/sql"
	"fmt"
	"suning/model"
)

func TakeOutGoodComment(goodName string) model.Comments {
	var id int
	var ids []int
	var c model.Comment
	var C model.Comments
	rows, err := dB.Query("select id,writer,comment,comment_time from comment where good=? && parent_id=0", goodName)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&id, &c.Writer, &c.Comment, &c.CommentTime)
		if err != nil {
			fmt.Println(err)
		}
		ids = append(ids, id)
	}
	ids, C = GetComment(ids, C)
	return C
}

func BringComments(comment, writer, good, t string) {
	_, err := dB.Exec("insert into comment (parent_id,comment,writer,good,comment_time) value (?,?,?,?,?)",
		0, comment, writer, good, t)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetComment(ids []int, C model.Comments) ([]int, model.Comments) {
	var id int
	var Id []int
	var c model.Comment
	var rows []*sql.Row
	var errs []error
	for _, i := range ids {
		row := dB.QueryRow("select id,writer,comment,respond,res_writer,respond_time from comment where parent_id=?", i)
		rows = append(rows, row)
	}
	for _, r := range rows {
		err := r.Scan(&id, &c.Writer, &c.Comment, &c.Respond, &c.ResWriter, &c.RespondTime)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		Id = append(Id, id)
		C = append(C, c)
	}
	if len(errs) != len(rows) {
		Id, C = GetComment(Id, C)
		errs = []error{}
	}
	return ids, C
}

func DeleteRes(username, comment, writerTo, respond, good, resTime string) {
	row := dB.QueryRow("delete from comment where res_writer=? && comment=? && writer=? && respond=? && good=? && respond_time=?",
		username, comment, writerTo, respond, good, resTime)
	err := row.Scan()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func PublishRespond(username, comment, writerTo, respond, t, T, good string) {
	var id int
	row := dB.QueryRow("select id from comment where writer=? && comment=? && comment_time=? && good=?",
		writerTo, comment, T, good)
	err := row.Scan(&id)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = dB.Exec("insert into comment (parent_id,res_writer,comment,writer,respond,respond_time,comment_time,good) value (?,?,?,?,?,?,?,?)",
		id, username, comment, writerTo, respond, t, T, good)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DeleteComment(username, comment, time, good string) {
	row := dB.QueryRow("delete from comment where parent_id=0 && writer=? && comment=? && comment_time=? && good=?",
		username, comment, time, good)
	err := row.Scan()
	if err != nil {
		fmt.Println(err)
		return
	}
}
