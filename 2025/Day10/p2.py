from gurobipy import GRB, Model

f = open("input.txt")

lines = f.readlines()

res = 0
model = Model("Day 10 p2")
model.Params.LogToConsole = 0

for line in lines:
    line = line.split(" ")
    vx = [int(x) for x in range(len(line)-2)]
    vars = [model.addVar(name="x"+str(x), vtype=GRB.INTEGER, lb=0)
            for x in range(len(vx))]
    model.setObjective(sum(vars), GRB.MINIMIZE)
    all_buttons = []
    for button in line[1:-1]:
        pushed = [int(x) for x in button[1:-1].split(",")]
        all_buttons.append(pushed)
    bounds = [int(x) for x in line[-1][1:-2].split(",")]
    for bound in range(len(bounds)):
        b = []
        for button in range(len(all_buttons)):
            if bound in all_buttons[button]:
                b.append(button)
        model.addConstr(sum([vars[x] for x in b]) == bounds[bound])

    model.optimize()
    res += model.ObjVal

print(res)
