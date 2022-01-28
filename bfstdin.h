#include <stdio.h>
#include <windows.h>
#include <stdlib.h>

// Function take password and
// reset to console mode
char inputPswrd()
{
    HANDLE hStdInput = GetStdHandle(STD_INPUT_HANDLE);
    DWORD mode = 0;

    // Create a restore point Mode
    // is know 503
    GetConsoleMode(hStdInput, &mode);

    // Enable echo input
    // set to 499
    SetConsoleMode(
        hStdInput,
        mode & (~ENABLE_ECHO_INPUT));

    // Take input
    char ipt;
    scanf("%c", &ipt);

    // Otherwise next cout will print
    // into the same line
    printf("\n");

    // Restore the mode
    SetConsoleMode(hStdInput, mode);

    return ipt;
}
