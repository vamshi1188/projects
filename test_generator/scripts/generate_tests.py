#!/usr/bin/env python3
"""
C++ Unit Test Generator

This script generates unit tests for C++ applications using a Large Language Model (LLM).
It analyzes C++ code, generates appropriate unit tests, refines them, and helps improve test coverage.
"""

import os
import sys
import argparse
import glob
import subprocess
import yaml
import json
import requests
from pathlib import Path
import logging

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger('test_generator')

# Default values
DEFAULT_MODEL = "gpt-4"
DEFAULT_OUTPUT_DIR = "tests"
DEFAULT_PROMPT_DIR = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), "prompts")

def parse_arguments():
    """Parse command line arguments."""
    parser = argparse.ArgumentParser(description='Generate unit tests for C++ code using LLM')
    parser.add_argument('source_dir', help='Directory containing C++ source files')
    parser.add_argument('--output', '-o', default=DEFAULT_OUTPUT_DIR, 
                        help=f'Output directory for generated tests (default: {DEFAULT_OUTPUT_DIR})')
    parser.add_argument('--model', '-m', default=DEFAULT_MODEL,
                        help=f'LLM model to use (default: {DEFAULT_MODEL})')
    parser.add_argument('--api-key', '-k', 
                        help='API key for LLM service (defaults to LLM_API_KEY environment variable)')
    parser.add_argument('--prompt-dir', '-p', default=DEFAULT_PROMPT_DIR,
                        help=f'Directory containing prompt YAML files (default: {DEFAULT_PROMPT_DIR})')
    parser.add_argument('--verbose', '-v', action='store_true',
                        help='Enable verbose output')
    
    args = parser.parse_args()
    
    # Set verbose logging if requested
    if args.verbose:
        logger.setLevel(logging.DEBUG)
    
    # Check if source directory exists
    if not os.path.isdir(args.source_dir):
        parser.error(f"Source directory '{args.source_dir}' does not exist")
    
    # Check if prompt directory exists
    if not os.path.isdir(args.prompt_dir):
        parser.error(f"Prompt directory '{args.prompt_dir}' does not exist")
    
    # Get API key from environment if not provided
    if not args.api_key:
        args.api_key = os.environ.get('LLM_API_KEY')
        if not args.api_key:
            parser.error("API key must be provided via --api-key or LLM_API_KEY environment variable")
    
    return args

def load_yaml_prompt(prompt_file):
    """Load a YAML prompt file."""
    try:
        with open(prompt_file, 'r') as f:
            return yaml.safe_load(f)
    except Exception as e:
        logger.error(f"Error loading prompt file {prompt_file}: {e}")
        sys.exit(1)

def find_cpp_files(source_dir):
    """Find all C++ source and header files in the given directory."""
    cpp_extensions = ['*.cpp', '*.cc', '*.h', '*.hpp']
    cpp_files = []
    
    for ext in cpp_extensions:
        cpp_files.extend(glob.glob(os.path.join(source_dir, '**', ext), recursive=True))
    
    return cpp_files

def read_file_content(file_path):
    """Read the content of a file."""
    try:
        with open(file_path, 'r') as f:
            return f.read()
    except Exception as e:
        logger.error(f"Error reading file {file_path}: {e}")
        return None

def call_llm_api(prompt, model, api_key):
    """Call the LLM API with the given prompt."""
    # This is a generic implementation that works with OpenAI's API
    # Modify this function if using a different LLM API
    
    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {api_key}"
    }
    
    data = {
        "model": model,
        "messages": [{"role": "user", "content": prompt}],
        "temperature": 0.7
    }
    
    try:
        response = requests.post(
            "https://api.openai.com/v1/chat/completions",
            headers=headers,
            json=data
        )
        response.raise_for_status()
        return response.json()["choices"][0]["message"]["content"]
    except Exception as e:
        logger.error(f"Error calling LLM API: {e}")
        if hasattr(e, 'response') and hasattr(e.response, 'text'):
            logger.error(f"API response: {e.response.text}")
        return None

def generate_initial_tests(cpp_files, prompt_data, model, api_key):
    """Generate initial unit tests for the given C++ files."""
    logger.info("Generating initial unit tests...")
    
    # Create a prompt with the C++ code
    code_content = ""
    for file_path in cpp_files:
        content = read_file_content(file_path)
        if content:
            code_content += f"\n// File: {file_path}\n{content}\n"
    
    # Replace placeholder in the prompt template
    prompt_template = prompt_data.get("prompt", "")
    prompt = prompt_template.replace("{{CODE}}", code_content)
    
    # Call the LLM API
    response = call_llm_api(prompt, model, api_key)
    if not response:
        logger.error("Failed to generate initial tests")
        return None
    
    return response

def refine_tests(initial_tests, prompt_data, model, api_key):
    """Refine the generated tests."""
    logger.info("Refining generated tests...")
    
    # Replace placeholder in the prompt template
    prompt_template = prompt_data.get("prompt", "")
    prompt = prompt_template.replace("{{TESTS}}", initial_tests)
    
    # Call the LLM API
    response = call_llm_api(prompt, model, api_key)
    if not response:
        logger.error("Failed to refine tests")
        return initial_tests  # Return the initial tests if refinement fails
    
    return response

def fix_build_issues(tests, build_output, prompt_data, model, api_key):
    """Fix build issues in the generated tests."""
    logger.info("Fixing build issues...")
    
    # Replace placeholders in the prompt template
    prompt_template = prompt_data.get("prompt", "")
    prompt = prompt_template.replace("{{TESTS}}", tests).replace("{{BUILD_OUTPUT}}", build_output)
    
    # Call the LLM API
    response = call_llm_api(prompt, model, api_key)
    if not response:
        logger.error("Failed to fix build issues")
        return tests  # Return the original tests if fixing fails
    
    return response

def improve_coverage(tests, coverage_output, prompt_data, model, api_key):
    """Improve test coverage based on coverage analysis."""
    logger.info("Improving test coverage...")
    
    # Replace placeholders in the prompt template
    prompt_template = prompt_data.get("prompt", "")
    prompt = prompt_template.replace("{{TESTS}}", tests).replace("{{COVERAGE_OUTPUT}}", coverage_output)
    
    # Call the LLM API
    response = call_llm_api(prompt, model, api_key)
    if not response:
        logger.error("Failed to improve coverage")
        return tests  # Return the original tests if improvement fails
    
    return response

def save_tests(tests, output_dir):
    """Save the generated tests to files."""
    os.makedirs(output_dir, exist_ok=True)
    
    # Extract test files from the LLM response
    # This is a simple implementation that assumes the LLM returns test files
    # with filenames indicated by "// Filename: test_xxx.cc" comments
    
    current_file = None
    current_content = []
    
    for line in tests.split('\n'):
        if line.startswith('// Filename:') or line.startswith('// File:'):
            # Save the previous file if it exists
            if current_file and current_content:
                file_path = os.path.join(output_dir, current_file)
                with open(file_path, 'w') as f:
                    f.write('\n'.join(current_content))
                logger.info(f"Saved test file: {file_path}")
                current_content = []
            
            # Extract the new filename
            current_file = line.split(':', 1)[1].strip()
        else:
            current_content.append(line)
    
    # Save the last file
    if current_file and current_content:
        file_path = os.path.join(output_dir, current_file)
        with open(file_path, 'w') as f:
            f.write('\n'.join(current_content))
        logger.info(f"Saved test file: {file_path}")

def build_tests(output_dir):
    """Build the generated tests and return the build output."""
    logger.info("Building tests...")
    
    try:
        # Run CMake to build the tests
        build_dir = os.path.join(output_dir, 'build')
        os.makedirs(build_dir, exist_ok=True)
        
        # Run CMake
        cmake_cmd = ['cmake', '..']
        cmake_process = subprocess.run(
            cmake_cmd, 
            cwd=build_dir, 
            capture_output=True, 
            text=True
        )
        
        # Run Make
        make_cmd = ['make']
        make_process = subprocess.run(
            make_cmd, 
            cwd=build_dir, 
            capture_output=True, 
            text=True
        )
        
        # Combine outputs
        build_output = f"CMake output:\n{cmake_process.stdout}\n{cmake_process.stderr}\n"
        build_output += f"Make output:\n{make_process.stdout}\n{make_process.stderr}"
        
        # Check if build was successful
        if cmake_process.returncode != 0 or make_process.returncode != 0:
            logger.error("Build failed")
            return build_output, False
        
        logger.info("Build successful")
        return build_output, True
        
    except Exception as e:
        logger.error(f"Error building tests: {e}")
        return str(e), False

def run_coverage_analysis(output_dir):
    """Run the tests with coverage analysis and return the coverage output."""
    logger.info("Running coverage analysis...")
    
    try:
        # Run the tests with coverage
        build_dir = os.path.join(output_dir, 'build')
        
        # Run the tests with gcov
        test_cmd = ['./org_chart_test', '--gtest_output=xml:test_results.xml']
        test_process = subprocess.run(
            test_cmd, 
            cwd=build_dir, 
            capture_output=True, 
            text=True
        )
        
        # Run lcov to generate coverage info
        lcov_cmd = ['lcov', '--capture', '--directory', '.', '--output-file', 'coverage.info']
        lcov_process = subprocess.run(
            lcov_cmd, 
            cwd=build_dir, 
            capture_output=True, 
            text=True
        )
        
        # Generate HTML report
        genhtml_cmd = ['genhtml', 'coverage.info', '--output-directory', 'coverage_report']
        genhtml_process = subprocess.run(
            genhtml_cmd, 
            cwd=build_dir, 
            capture_output=True, 
            text=True
        )
        
        # Combine outputs
        coverage_output = f"Test output:\n{test_process.stdout}\n{test_process.stderr}\n"
        coverage_output += f"LCOV output:\n{lcov_process.stdout}\n{lcov_process.stderr}\n"
        coverage_output += f"GenHTML output:\n{genhtml_process.stdout}\n{genhtml_process.stderr}"
        
        logger.info("Coverage analysis completed")
        return coverage_output
        
    except Exception as e:
        logger.error(f"Error running coverage analysis: {e}")
        return str(e)

def main():
    """Main function."""
    args = parse_arguments()
    
    # Find C++ files
    cpp_files = find_cpp_files(args.source_dir)
    if not cpp_files:
        logger.error(f"No C++ files found in {args.source_dir}")
        sys.exit(1)
    
    logger.info(f"Found {len(cpp_files)} C++ files")
    
    # Load prompt YAML files
    initial_prompt = load_yaml_prompt(os.path.join(args.prompt_dir, 'initial.yaml'))
    refine_prompt = load_yaml_prompt(os.path.join(args.prompt_dir, 'refine.yaml'))
    build_prompt = load_yaml_prompt(os.path.join(args.prompt_dir, 'build.yaml'))
    coverage_prompt = load_yaml_prompt(os.path.join(args.prompt_dir, 'coverage.yaml'))
    
    # Generate initial tests
    initial_tests = generate_initial_tests(cpp_files, initial_prompt, args.model, args.api_key)
    if not initial_tests:
        sys.exit(1)
    
    # Refine tests
    refined_tests = refine_tests(initial_tests, refine_prompt, args.model, args.api_key)
    
    # Save refined tests
    save_tests(refined_tests, args.output)
    
    # Build tests
    build_output, build_success = build_tests(args.output)
    
    # Fix build issues if needed
    if not build_success:
        fixed_tests = fix_build_issues(refined_tests, build_output, build_prompt, args.model, args.api_key)
        save_tests(fixed_tests, args.output)
        
        # Try building again
        build_output, build_success = build_tests(args.output)
        if not build_success:
            logger.error("Failed to fix build issues")
            sys.exit(1)
    
    # Run coverage analysis
    coverage_output = run_coverage_analysis(args.output)
    
    # Improve coverage
    improved_tests = improve_coverage(refined_tests, coverage_output, coverage_prompt, args.model, args.api_key)
    save_tests(improved_tests, args.output)
    
    logger.info("Test generation completed successfully")

if __name__ == "__main__":
    main()