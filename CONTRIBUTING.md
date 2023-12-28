# Contributing to OPREC-NETLAB-BE

Welcome  friends! We're excited that you're interested in contributing. Please take a moment to read the following guidelines.

## Getting Started

1. **Clone the Repository:**
   - Clone the repository to your local machine:
     ```
     git clone https://github.com/bernanda0/oprec-netlab-be.git
     cd oprec-netlab-be
     ```

2. **Create a Branch:**
   - Create a branch for your work:
     ```
     git checkout -b feature/your-feature-name
     ```

3. **If it's your first time contributinng, then kindly setup the .env**
    - Rename the `.env.template` to `.env`
    - Fill each variables with its appropriate value based on your local settings/configuration.

4. **Make Changes:**
   - Implement your changes and ensure they adhere to the project's coding standards.

5. **Commit Changes:**
   - Commit your changes with a clear and concise message:
     ```
     git commit -m "Implement feature: your-feature-name"
     ```

6. **Push Changes:**
   - Push your changes :
     ```
     git push origin feature/your-feature-name
     ```

7. **Open a Pull Request:**
   - Open a pull request from your branch to the `master` branch of the main repository.
   - Provide a detailed description of your changes and reference any related issues.

## Branching Strategy

We follow the GitHub flow for branching:

- **`master`:** The development branch. Features are merged into this branch for testing.
  
- **`stable`:** The stable/production branch. Always deployable and reflects the current production release.

- **`feature/`:** Create feature branches for new features or enhancements.
  - Example: `feature/new-feature`

- **`bugfix/`:** Create bugfix branches for fixing issues.
  - Example: `bugfix/fix-issue`

## Code Quality

- **Build and Compile Tests:**
  - Before submitting a pull request, ensure that your changes pass the build and compile tests.
  - Use the following commands to check:
    ```
    make build
    make test
    ```
  - Your changes must pass these tests to maintain the integrity of the project.

## Code Reviews

- All contributions will go through a code review process.
- Address feedback and make necessary changes based on reviewer comments.

