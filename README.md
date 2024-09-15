![Project Diagram](images/bottle.png "Project Diagram")
To rewrite the provided instructions for a README.md format, which is often used on GitHub or other version control platforms to explain setup and usage, here’s a structured and clear version that details how to download, install, and run the `Bottle.pkg`, and how to troubleshoot potential path issues specifically for a macOS environment using `zsh` as the shell.

### Installation and Running Instructions for Bottle on macOS

#### Step 1: Download and Install `Bottle.pkg`
1. **Download the Package**: Download `Bottle.pkg` from the provided source or repository.
2. **Install the Package**: Double-click on `Bottle.pkg` to start the installation process and follow the on-screen instructions to complete the installation.

#### Step 2: Running the Bottle Command
1. **Open Terminal**: You can open Terminal by pressing `Cmd + Space` to open Spotlight, type "Terminal", and hit Enter.
2. **Run the Command**:
   ```bash
   bottle <filename>.bottle
   ```
   Replace `<filename>.bottle` with the actual name of the file you want to process.

#### Step 3: Troubleshooting
If you receive a "command not found" error when trying to run `bottle`, follow these steps to resolve it:

##### Ensure `bottle` is in Your PATH
1. **Open Spotlight**:
   - Press `Cmd + Space` to open Spotlight.
   - Type `Monke-1.0.0.pkg` to locate the installation directory. (Note: Replace "Monke-1.0.0.pkg" with "Bottle.pkg" if searching for Bottle's installation directory).
   - Press `Cmd + Enter` (or `Cmd + Click`) to open the enclosing folder in Finder.

2. **Update Your PATH**:
   - Open Terminal and run the following command to add the correct directory to your PATH in `~/.zshrc`:
     ```bash
     echo 'export PATH="/Users/utsavdey/Desktop/Monke/MyPackageRoot/usr/local/bin:$PATH"' >> ~/.zshrc
     ```
   - Make sure to adjust the path to the directory where `bottle` is located, not including the executable itself.
   - Reload your `~/.zshrc` to apply the changes:
     ```bash
     source ~/.zshrc
     ```

3. **Verify Installation**:
   - Check that `bottle` is indeed in the specified directory:
     ```bash
     ls -l /Users/utsavdey/Desktop/Monke/MyPackageRoot/usr/local/bin/bottle
     ```
   - Ensure `bottle` is executable:
     ```bash
     chmod +x /Users/utsavdey/Desktop/Monke/MyPackageRoot/usr/local/bin/bottle
     ```

#### Additional Notes
- If you still encounter issues, ensure you have administrative rights to change environment settings or install software, as required.
- For further assistance, consult the documentation provided with `Bottle.pkg` or reach out to support.

This structured format for a README.md provides clear instructions on how to set up, run, and troubleshoot the `bottle` command, making it easier for users to follow and reduce setup errors.
