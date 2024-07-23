## Suggestions for the code
1. Remove the unused variables, such as `f1`, `err`, `fs`.
2. add error handling for function `FetchAllFoldersByOrgID`, this should be checked after we use it and before other steps.
3. Delete the loops for transfers types of Folders or pointers, folder pointers are already stored in `r`.