# C++ Unit Test Generator

## Overview
This tool automatically generates unit tests for C++ applications using a Large Language Model (LLM). It analyzes your C++ code, generates appropriate unit tests, refines them based on feedback, and helps improve test coverage.

## Architecture
The test generator follows a multi-stage workflow:
1. **Initial Test Generation**: Analyzes C++ code and generates initial unit tests
2. **Test Refinement**: Improves tests by removing duplicates and enhancing coverage
3. **Build & Debug**: Handles build issues and improves tests based on compiler feedback
4. **Coverage Analysis**: Analyzes test coverage and suggests improvements

## Requirements
- Python 3.7+
- Access to an LLM API (OpenAI, Anthropic, etc.)
- C++ compiler (GCC or Clang)
- Google Test framework
- GNU code coverage tools (gcov, lcov)

## Installation
No local LLM installation is required. The tool uses remote LLM APIs.

```bash
# Clone the repository
git clone <repository-url>

# Install Python dependencies
pip install -r requirements.txt

# Configure your API key
export LLM_API_KEY=your_api_key_here
```

## Usage

### Basic Usage
```bash
./test_generator/scripts/run_test_generator.sh <source_directory>
```

### Advanced Usage
```bash
./test_generator/scripts/run_test_generator.sh \
  --source <source_directory> \
  --output <output_directory> \
  --model <model_name> \
  --api-key <api_key>
```

## Customization
You can customize the test generation process by modifying the YAML files in the `test_generator/prompts` directory:
- `initial.yaml`: Instructions for initial test generation
- `refine.yaml`: Instructions for test refinement
- `build.yaml`: Instructions for fixing build issues
- `coverage.yaml`: Instructions for improving test coverage

## Example

### Input
```cpp
// Example C++ code
class Calculator {
public:
    int add(int a, int b) {
        return a + b;
    }
};
```

### Output
```cpp
// Generated test
TEST(CalculatorTest, AdditionWorks) {
    Calculator calc;
    EXPECT_EQ(calc.add(2, 3), 5);
    EXPECT_EQ(calc.add(-2, 2), 0);
    EXPECT_EQ(calc.add(0, 0), 0);
}
```

## Limitations
- The quality of generated tests depends on the LLM being used
- Complex code may require manual refinement of generated tests
- Some edge cases might not be automatically covered

## License
This project is licensed under the MIT License - see the LICENSE file for details.