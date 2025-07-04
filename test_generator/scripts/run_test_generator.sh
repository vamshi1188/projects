#!/bin/bash
# C++ Unit Test Generator Runner Script
# This script runs the test generator Python script with the provided arguments

set -e

# Get the directory of this script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$( cd "$SCRIPT_DIR/../.." && pwd )"
GENERATOR_DIR="$( cd "$SCRIPT_DIR/.." && pwd )"

# Default values
SOURCE_DIR="$PROJECT_ROOT"
OUTPUT_DIR="$PROJECT_ROOT/tests"
MODEL="gpt-4"
PROMPT_DIR="$GENERATOR_DIR/prompts"
API_KEY=${LLM_API_KEY:-""}
VERBOSE=0

# Function to display usage information
function show_usage {
    echo "Usage: $0 [options] [source_directory]"
    echo ""
    echo "Options:"
    echo "  -h, --help                 Show this help message and exit"
    echo "  -o, --output DIR           Output directory for generated tests (default: $OUTPUT_DIR)"
    echo "  -m, --model MODEL          LLM model to use (default: $MODEL)"
    echo "  -k, --api-key KEY          API key for LLM service (defaults to LLM_API_KEY environment variable)"
    echo "  -p, --prompt-dir DIR       Directory containing prompt YAML files (default: $PROMPT_DIR)"
    echo "  -v, --verbose              Enable verbose output"
    echo ""
    echo "Arguments:"
    echo "  source_directory           Directory containing C++ source files (default: $SOURCE_DIR)"
    echo ""
    echo "Example:"
    echo "  $0 --output ./my_tests --model gpt-3.5-turbo ./src"
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_usage
            exit 0
            ;;
        -o|--output)
            OUTPUT_DIR="$2"
            shift 2
            ;;
        -m|--model)
            MODEL="$2"
            shift 2
            ;;
        -k|--api-key)
            API_KEY="$2"
            shift 2
            ;;
        -p|--prompt-dir)
            PROMPT_DIR="$2"
            shift 2
            ;;
        -v|--verbose)
            VERBOSE=1
            shift
            ;;
        *)
            # Assume it's the source directory
            SOURCE_DIR="$1"
            shift
            ;;
    esac
done

# Check if Python is installed
if ! command -v python3 &> /dev/null; then
    echo "Error: Python 3 is required but not installed."
    exit 1
fi

# Check if required Python packages are installed
REQUIRED_PACKAGES=("pyyaml" "requests")
MISSING_PACKAGES=()

for package in "${REQUIRED_PACKAGES[@]}"; do
    if ! python3 -c "import $package" &> /dev/null; then
        MISSING_PACKAGES+=("$package")
    fi
done

if [ ${#MISSING_PACKAGES[@]} -gt 0 ]; then
    echo "Error: The following Python packages are required but not installed:"
    for package in "${MISSING_PACKAGES[@]}"; do
        echo "  - $package"
    done
    echo "Please install them using: pip install ${MISSING_PACKAGES[*]}"
    exit 1
fi

# Check if API key is provided
if [ -z "$API_KEY" ]; then
    echo "Error: API key is required. Please provide it using --api-key option or set the LLM_API_KEY environment variable."
    exit 1
fi

# Check if prompt directory exists and contains required files
if [ ! -d "$PROMPT_DIR" ]; then
    echo "Creating prompt directory: $PROMPT_DIR"
    mkdir -p "$PROMPT_DIR"
fi

REQUIRED_PROMPTS=("initial.yaml" "refine.yaml" "build.yaml" "coverage.yaml")
MISSING_PROMPTS=()

for prompt in "${REQUIRED_PROMPTS[@]}"; do
    if [ ! -f "$PROMPT_DIR/$prompt" ]; then
        MISSING_PROMPTS+=("$prompt")
    fi
done

if [ ${#MISSING_PROMPTS[@]} -gt 0 ]; then
    echo "Warning: The following prompt files are missing:"
    for prompt in "${MISSING_PROMPTS[@]}"; do
        echo "  - $prompt"
    done
    echo "Creating default prompt files..."
    
    # Create default prompt files
    for prompt in "${MISSING_PROMPTS[@]}"; do
        case $prompt in
            initial.yaml)
                cat > "$PROMPT_DIR/$prompt" << 'EOF'
prompt: |
  You are an expert C++ developer tasked with creating unit tests for the following C++ code.
  Please generate comprehensive unit tests using Google Test framework.
  
  Here's the code to test:
  
  {{CODE}}
  
  Please generate unit tests that:
  1. Test all public methods and functions
  2. Include edge cases and error conditions
  3. Have good test coverage
  4. Follow best practices for Google Test
  
  Format your response as complete test files, with each file starting with a comment line:
  // Filename: test_filename.cc
EOF
                ;;
            refine.yaml)
                cat > "$PROMPT_DIR/$prompt" << 'EOF'
prompt: |
  You are an expert C++ developer tasked with refining the following unit tests.
  Please improve these tests by:
  
  1. Removing any duplicate tests
  2. Adding missing includes and dependencies
  3. Ensuring proper test fixture setup and teardown
  4. Improving test coverage for edge cases
  5. Following Google Test best practices
  
  Here are the tests to refine:
  
  {{TESTS}}
  
  Format your response as complete test files, with each file starting with a comment line:
  // Filename: test_filename.cc
EOF
                ;;
            build.yaml)
                cat > "$PROMPT_DIR/$prompt" << 'EOF'
prompt: |
  You are an expert C++ developer tasked with fixing build issues in the following unit tests.
  
  Here are the tests that failed to build:
  
  {{TESTS}}
  
  And here is the build output showing the errors:
  
  {{BUILD_OUTPUT}}
  
  Please fix all build issues and return the corrected tests.
  
  Format your response as complete test files, with each file starting with a comment line:
  // Filename: test_filename.cc
EOF
                ;;
            coverage.yaml)
                cat > "$PROMPT_DIR/$prompt" << 'EOF'
prompt: |
  You are an expert C++ developer tasked with improving test coverage for the following unit tests.
  
  Here are the current tests:
  
  {{TESTS}}
  
  And here is the coverage analysis:
  
  {{COVERAGE_OUTPUT}}
  
  Please improve the tests to increase coverage by:
  1. Adding tests for uncovered functions and methods
  2. Adding tests for uncovered branches and conditions
  3. Ensuring all edge cases are tested
  
  Format your response as complete test files, with each file starting with a comment line:
  // Filename: test_filename.cc
EOF
                ;;
        esac
    done
fi

# Run the Python script
echo "Running test generator..."
echo "Source directory: $SOURCE_DIR"
echo "Output directory: $OUTPUT_DIR"
echo "Model: $MODEL"
echo "Prompt directory: $PROMPT_DIR"

VERBOSE_FLAG=""
if [ $VERBOSE -eq 1 ]; then
    VERBOSE_FLAG="--verbose"
fi

python3 "$SCRIPT_DIR/generate_tests.py" \
    --output "$OUTPUT_DIR" \
    --model "$MODEL" \
    --api-key "$API_KEY" \
    --prompt-dir "$PROMPT_DIR" \
    $VERBOSE_FLAG \
    "$SOURCE_DIR"

echo "Test generation completed successfully!"
echo "Generated tests are in: $OUTPUT_DIR"