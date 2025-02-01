#!/bin/bash

# Check if the problem number is passed as an argument
if [ -z "$1" ]; then
  echo "Please provide a problem number and problem name (e.g., ./create_problem.sh '1. Two Sum')"
  exit 1
fi

# Problem number passed as the first argument
problem_number=$(echo "$1" | cut -d'.' -f1)  # "1. Two Sum" -> "1"
problem_title=$1 # "1. Two Sum"
formatted_number=$(printf "%04d" "$problem_number")  # "1" -> "0001"

# Base directory where problems will be stored
base_dir="Problems"
problem_dir="${base_dir}/${formatted_number}"

# Create the directory if it doesn't exist
mkdir -p "$problem_dir"

# solution.go faylini yaratish
solution_file="$problem_dir/solution.go"
echo "package problem$formatted_number" > "$solution_file"

# solution_test.go faylini yaratish
test_file="$problem_dir/solution_test.go"
cat <<EOL > "$test_file"
package problem${formatted_number}_test

import (
	"testing"

	problem${formatted_number} "github.com/realtemirov/leetcode/Problems/${formatted_number}"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected int
	}{
		{
			name:     "Test 1",
			cases:    []int{},
			expected: 0,
		},
		{
			name:     "Test 2",
			cases:    []int{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem${formatted_number}.Solution(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
EOL

# notes.md faylini yaratish
notes_file="$problem_dir/$problem_title.md"
cat <<EOL > "$notes_file"
# $problem_title

🟩 Easy | 🟧 Medium | 🟥 Hard

## Solution

My Solution

\`\`\`go
\`\`\`

![result]($problem_number.png)

Leetcode: [link]()
EOL

# Print success message
echo "Folder and files for problem $problem_number created successfully in $problem_dir"

# # Open VS Code
# cd "$problem_dir"
# code "$problem_title.md"
# code "solution.go"
# code "solution_test.go"
# cd ../

# Adding to README.md
readme_file="Problems/README.md"
cat <<EOL >> "$readme_file"
* 🟩 Easy | 🟧 Medium | 🟥 Hard - [${problem_title}](<./${formatted_number}/${problem_title}.md>)
EOL

# Print success message
echo "README.md updated successfully"

# Adding to SUMMARY.md
summary_file="SUMMARY.md"
cat <<EOL >> "$summary_file"
* [${problem_title}](<./Problems/${formatted_number}/${problem_title}.md>)
EOL

# Print success message
echo "SUMMARY.md updated successfully"