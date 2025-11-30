import sys

def write_dot(input_file, output_file):
    with open(output_file, 'w') as dot_file:
        dot_file.write("""digraph G {
    rankdir=LR;
    node [fontname="Arial"];
    edge [fontname="Arial"];

    node [style="filled", fillcolor="lightgray", shape="ellipse"];
    edge [color="#666666"];

""")

        operation_colors = {
            'AND': '#FFD700',
            'OR': '#98FB98',
            'XOR': '#87CEEB'
        }

        with open(input_file, 'r') as input:
            for line in input:
                line = line.strip()
                if not line:
                    continue

                parts = line.split('->')
                if len(parts) != 2:
                    continue

                inputs, output = parts[0].strip(), parts[1].strip()

                operation = None
                for op in ['AND', 'XOR', 'OR']:
                    if op in inputs:
                        operation = op
                        input_vars = inputs.split(op)
                        break

                if not operation:
                    continue

                input_vars = [var.strip() for var in input_vars]
                output = output.strip()

                operation_node = f"op_{input_vars[0]}_{input_vars[1]}_{operation}"

                dot_file.write(f'    "{input_vars[0]}" [fillcolor="white"]\n')
                dot_file.write(f'    "{input_vars[1]}" [fillcolor="white"]\n')
                dot_file.write(f'    "{output}" [fillcolor="#E8E8E8"]\n')
                dot_file.write(f'    "{operation_node}" [shape=box, fillcolor="{operation_colors[operation]}", '
                               f'label="{operation}", style="filled,rounded"]\n')

                dot_file.write(f'    "{input_vars[0]}" -> "{operation_node}" [arrowsize=0.7]\n')
                dot_file.write(f'    "{input_vars[1]}" -> "{operation_node}" [arrowsize=0.7]\n')
                dot_file.write(f'    "{operation_node}" -> "{output}" [arrowsize=0.7]\n')

        dot_file.write("}\n")


input_file = sys.argv[1]
output_file = sys.argv[2]
write_dot(input_file, output_file)
