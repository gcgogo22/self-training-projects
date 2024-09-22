# How to use Git Add with Folder Path or Filename Pattern 

In Git, the `git add` command allows you to stage changes for commit, and you can specify certain folders or file name patterns to only stage the changes you want.

Here’s how you can use `git add` to stage changes from a **specific folder** or **file name pattern**:

### 1. **Staging Changes from a Specific Folder**

To stage all changes (including new, modified, and deleted files) from a specific folder, you can use:

```bash
git add <folder_path>
```

For example, if you want to stage all changes inside the `src/` folder:

```bash
git add src/
```

This will add all changes (tracked and untracked files) inside the `src/` folder.

### 2. **Staging Files Matching a Certain File Name Pattern**

If you want to stage changes based on a specific file name pattern (e.g., all `.txt` files or all files starting with `config`), you can use file name patterns like:

```bash
git add *.txt
```

This will stage all `.txt` files in the current directory and its subdirectories.

For example, to stage all files with a `.go` extension in a specific folder:

```bash
git add src/*.go
```

This will stage all `.go` files in the `src/` folder.

### 3. **Staging Changes Recursively for All Matching Files**

If you want to recursively stage files that match a certain pattern from all directories (e.g., all `.js` files in all subdirectories):

```bash
git add '**/*.js'
```

Note that the `**` wildcard matches files in the current directory and all subdirectories.

### 4. **Staging Untracked Files Only**

To stage only **untracked files** (new files that Git hasn’t started tracking yet) from a specific folder:

```bash
git add <folder_path> --intent-to-add
```

Or to stage all untracked files that match a pattern:

```bash
git add '*.log' --intent-to-add
```

### 5. **Examples**

#### Example 1: Staging All Files in a Folder
To stage all changes in the `app/` folder:

```bash
git add app/
```

#### Example 2: Staging All `.md` Files
To stage all markdown files in the current directory:

```bash
git add *.md
```

#### Example 3: Staging All Files Starting with `config`
To stage all files that start with `config` in the current directory:

```bash
git add config*
```

#### Example 4: Staging Changes in a Folder with a Pattern
To stage all `.html` files in the `website/` folder:

```bash
git add website/*.html
```

### Summary:
- **Staging a folder**: `git add <folder_path>`
- **Staging by file pattern**: `git add *.txt` or `git add '**/*.js'`
- **Recursively staging by pattern**: `git add '**/*.pattern'`
- **Untracked files only**: `git add <pattern> --intent-to-add`

This approach helps you stage only the changes you want, based on file paths or patterns, making your commits more controlled and specific.