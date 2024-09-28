#!/bin/bash

# Check if the problem number is passed as an argument
if [ -z "$1" ]; then
  echo "Please provide a problem number (e.g., ./create_problem.sh 0002)"
  exit 1
fi

# Problem number passed as the first argument
problem_number=$1

# Base directory where problems will be stored
base_dir="Problems"

# Full path to the problem directory
problem_dir="$base_dir/$problem_number"

# Create the directory if it doesn't exist
mkdir -p "$problem_dir"

# Create solution.go with package name
solution_file="$problem_dir/solution.go"
echo "package problem$problem_number" > "$solution_file"

# Create solution_test.go with package name and a basic test function
test_file="$problem_dir/solution_test.go"
test_file="$problem_dir/solution_test.go"
cat <<EOL > "$test_file"
package problem${problem_number}_test

import (
	"testing"

	problem${problem_number} "github.com/realtemirov/leetcode/Problems/${problem_number}"
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
			expected: 4,
		},
		{
			name:     "Test 3",
			cases:    []int{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem${problem_number}.MaxProfit(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
EOL

# Create notes.md file
notes_file="$problem_dir/$problem_number.md"
cat <<EOL > "$notes_file"
# Problem $problem_number
🟩 Easy | 🟧 Medium | 🟥 Hard

## Example 1:
> **Input**:  \\
> **Output**:  \\
> **Explanation**: 

## Example 2:
> **Input**:  \\
> **Output**:  \\
> **Explanation**: 

## Constraints:
* \`-10^5 <= n <= 10^5\`

## Solution
> **My Solution**
> \`\`\`go
> \`\`\`

![result]($problem_number.png)

Leetcode: [link]()
EOL
touch "$notes_file"

# Print success message
echo "Folder and files for problem $problem_number created successfully in $problem_dir"
