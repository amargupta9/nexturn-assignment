students = [("amar",90),("rohit",78),("kutta",7),("arti",89)]

filter_by_marks = list(filter( lambda student: student[1] >= 50 , students))

print("filter by marks:\n")
for name , mark in filter_by_marks:
   print(f"{name} : {mark}")