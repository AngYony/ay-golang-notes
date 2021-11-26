package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		//如果year的值是否非法的，返回一个错误
		return errors.New("年份错误")
	}
	d.year = year
	return nil //返回nil作为错误
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("月份错误")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("day值错误")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}
