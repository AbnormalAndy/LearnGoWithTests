# Days in Month

def is_leap(year):
    if year % 4 == 0:
        if year % 100 == 0:
            if year % 400 == 0:
                print("Leap Year")
            else:
                print("Not a Leap Year")
        else:
            print("Leap Year")
    else:
        print("Not a Leap Year")

def days_in_month():
    month_days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]


## Instructions

Convert the is_leap() Function

Convert the function so it returns TRUE or FALSE instead of printing.

Create a New Function Called days_in_month()

Modify the function days_in_month(), which should take a year and a month as inputs.

E.G. days_in_month(year=2022, month=2)

It will use this information to work out if the year is a leap year and decide the number of days in the month then return that output.

year = int(input())
month = int(input())
days = days_in_month(year, month)

To-Do:
- Convert is_leap() to Golang
- Determine Leap Year
- Print to Returns


