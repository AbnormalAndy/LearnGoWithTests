# Grading Program

Instructions

There is a database of `student_scores` in the format of a dictionary. The keys are the student names and the values are their exam scores.

Write a program that converts their scores to grades. By the end of the program, there should be a new dictionary called `student_grades` that should contain student names for keys and their grades for values.

The scoring criteria:
- Scores 91-100 Grades = "Outstanding"
- Scores 81-90 Grades = "Exceeds Expectations"
- Scores 71-80 Grades = "Acceptable"
- Scores 70 or Lower Grades = "Fail"

`student_scores = {
    "Harry": 81,
    "Ron": 78,
    "Hermione": 99,
    "Draco": 74,
    "Neville": 62,
}`

`student_grades = {
    "Harry": "Exceeds Expectations",
    "Ron": "Acceptable",
    "Hermione": "Outstanding",
    "Draco": "Acceptable",
    "Neville": "Fail",
}`


