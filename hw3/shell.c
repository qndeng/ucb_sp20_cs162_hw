#include <ctype.h>
#include <errno.h>
#include <fcntl.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <signal.h>
#include <sys/wait.h>
#include <termios.h>
#include <unistd.h>

#include "tokenizer.h"

/* Convenience macro to silence compiler warnings about unused function parameters. */
#define unused __attribute__((unused))

/* Whether the shell is connected to an actual terminal or not. */
bool shell_is_interactive;

/* File descriptor for the shell input */
int shell_terminal;

/* Terminal mode settings for the shell */
struct termios shell_tmodes;

/* Process group id for the shell */
pid_t shell_pgid;

int cmd_exit(struct tokens *tokens);
int cmd_help(struct tokens *tokens);
int cmd_pwd(struct tokens *tokens);
int cmd_cd(struct tokens *tokens);

/* Built-in command functions take token array (see parse.h) and return int */
typedef int cmd_fun_t(struct tokens *tokens);

/* Built-in command struct and lookup table */
typedef struct fun_desc {
  cmd_fun_t *fun;
  char *cmd;
  char *doc;
} fun_desc_t;

fun_desc_t cmd_table[] = {
  {cmd_help, "?", "show this help menu"},
  {cmd_exit, "exit", "exit the command shell"},
  {cmd_pwd, "pwd", "print the current working directory to std"},
  {cmd_cd, "cd", "takes one argument, a directory path, and changes the current working directory to that directory"},
};

/* Print the current working directory to std */
int cmd_pwd(unused struct tokens *tokens) {
  char buffer[80];
  getcwd(buffer, sizeof(buffer));
  fprintf(stdout, "%s\n", buffer);
  return 1;
}

/* Takes one argument, a directory path, and changes the current working directory to that directory */
int cmd_cd(unused struct tokens *tokens) {
  char* newdir = tokens_get_token(tokens, 1);
  chdir(newdir);
  return 1;
}

/* Prints a helpful description for the given command */
int cmd_help(unused struct tokens *tokens) {
  for (unsigned int i = 0; i < sizeof(cmd_table) / sizeof(fun_desc_t); i++)
    printf("%s - %s\n", cmd_table[i].cmd, cmd_table[i].doc);
  return 1;
}

/* Exits this shell */
int cmd_exit(unused struct tokens *tokens) {
  exit(0);
}

/* Looks up the built-in command, if it exists. */
int lookup(char cmd[]) {
  for (unsigned int i = 0; i < sizeof(cmd_table) / sizeof(fun_desc_t); i++)
    if (cmd && (strcmp(cmd_table[i].cmd, cmd) == 0))
      return i;
  return -1;
}

/* Intialization procedures for this shell */
void init_shell() {
  /* Our shell is connected to standard input. */
  shell_terminal = STDIN_FILENO;

  /* Check if we are running interactively */
  shell_is_interactive = isatty(shell_terminal);

  if (shell_is_interactive) {
    /* If the shell is not currently in the foreground, we must pause the shell until it becomes a
     * foreground process. We use SIGTTIN to pause the shell. When the shell gets moved to the
     * foreground, we'll receive a SIGCONT. */
    while (tcgetpgrp(shell_terminal) != (shell_pgid = getpgrp()))
      kill(-shell_pgid, SIGTTIN);

    /* Saves the shell's process id */
    shell_pgid = getpid();

    /* Take control of the terminal */
    tcsetpgrp(shell_terminal, shell_pgid);

    /* Save the current termios to a variable, so it can be restored later. */
    tcgetattr(shell_terminal, &shell_tmodes);
  }
}

int isFileExistsAccess(const char *path)
{
    // Check for file existence
    if (access(path, F_OK) == -1)
        return 0;

    return 1;
}




void execute(char *line){

	setpgid(getpid(), getpid());
    tcsetpgrp(0, getpid());
    signal(SIGINT, SIG_DFL);
    signal(SIGQUIT, SIG_DFL);
    signal(SIGTSTP, SIG_DFL);
    signal(SIGCONT, SIG_DFL);
    signal(SIGTTIN, SIG_DFL);
    signal(SIGTTOU, SIG_DFL);

    printf("cat pid: %d, cat pgid: %d cat foreground pgid: %d\n", getpid(), getpgid(getpid()), tcgetpgrp(0));

	// get args and path
	struct tokens *tokens = tokenize(line);
	unused char *path = tokens_get_token(tokens, 0);
	int len = tokens_get_length(tokens);
	unused char *args[len + 1];
	for(int i = 0;i < len;i++){
		args[i] = tokens_get_token(tokens, i);
	}
	args[len] = NULL;

	bool is_full_path = false;
	for(int i = 0;i < sizeof(path)/sizeof(char);i++){
		if(path[i] == '/'){
			is_full_path = true;
			break;
		}
	}
	if(!is_full_path){
		char *envs_tmp = getenv("PATH");
		char envs[999];
		strcpy(envs, envs_tmp);
		for(int i = 0;i < sizeof(envs)/sizeof(char);i++){
    		if(envs[i] == ':'){
    			envs[i] = ' ';
    		}
		}
		struct tokens* path_tokens = tokenize(envs);
		int path_len = tokens_get_length(path_tokens);
		bool find = false;
		for(int i = 0;i < path_len && !find;i++){
			char *dir = tokens_get_token(path_tokens, i);
			char slash[] = {'/', '\0'};
			char *possible_path =strcat(dir, strcat(slash, path));
			if(isFileExistsAccess(possible_path)){
				find = true;
				path = possible_path;
			}
		}
		if(!find){
			printf("%s\n", "cmd not exist");
			exit(1);
		} 
	}


	




	// function for redirection ( '<' , '>' )
    int in=0, out=0;
    char input[64], output[64];
    // finds where '<' or '>' occurs and make that argv[i] = NULL , to ensure that command wont't read that
    for(int i = 0;args[i] != NULL;i++)
    {
        if(strcmp(args[i],"<") == 0)
        {        
            args[i]=NULL;
            strcpy(input,args[i+1]);
            // args[i+1]=NULL;
            in=2;           
        }               

        if(strcmp(args[i],">") == 0)
        {      
            args[i]=NULL;
            strcpy(output,args[i+1]);
            // args[i+1]=NULL;
            out=2;
        }         
    }
    //if '<' char was found in string inputted by user
    if(in)
    {   
        // fdo is file-descriptor
        int fd0;
        if ((fd0 = open(input, O_RDONLY, 0)) < 0) {
            perror("Couldn't open input file");
            exit(0);
        }           
        // dup2() copies content of fdo in input of preceeding file
        dup2(fd0, STDIN_FILENO); // STDIN_FILENO here can be replaced by 0 

        close(fd0); // necessary
    }
    //if '>' char was found in string inputted by user 
    if (out)
    {
        int fd1 ;
        if ((fd1 = creat(output , 0644)) < 0) {
            perror("Couldn't open the output file");
            exit(0);
        }           

        dup2(fd1, STDOUT_FILENO); // 1 here can be replaced by STDOUT_FILENO
        close(fd1);
    }
	execv(path, args);
}


void executePrograms(char *line){
	int i = 0;
	int sep = 0;
	while(line[i] != '\0'){
		if(line[i] == '|'){
			sep = i;
			break;
		}
		i++;
	}

	if(sep == 0){
		execute(line);
	}else{
		char before[4096];
		char after[4096];
		strncpy(before, line, sep-1);
		before[sep-1] = '\0';
		strcpy(after, line+i+2);

		int pipefd[2];
		pid_t p1;
		if(pipe(pipefd) < 0) perror("pipe fail");
		p1 = fork();
		if(p1 < 0) perror("fork fail");
		if(p1){
			close(pipefd[1]);
			dup2(pipefd[0], STDIN_FILENO);
			executePrograms(after);
		}else{
			close(pipefd[0]);
			dup2(pipefd[1], STDOUT_FILENO);
			execute(before);
		}


	}
}


int main(unused int argc, unused char *argv[]) {
  init_shell();

  setpgid(getpid(), getpid());
  tcsetpgrp(0, getpid());

  signal(SIGINT, SIG_IGN);
  signal(SIGQUIT, SIG_IGN);
  signal(SIGTSTP, SIG_IGN);
  signal(SIGCONT, SIG_IGN);
  signal(SIGTTIN, SIG_IGN);
  signal(SIGTTOU, SIG_IGN);

  printf("shell pid: %d, shell pgid: %d shell foreground pgid: %d\n", getpid(), getpgid(getpid()), tcgetpgrp(0));

  static char line[4096];
  int line_num = 0;

  /* Please only print shell prompts when standard input is not a tty */
  if (shell_is_interactive)
    fprintf(stdout, "%d: ", line_num);

  while (fgets(line, 4096, stdin)) {
    /* Split our line into words. */
    struct tokens *tokens = tokenize(line);

    /* Find which built-in function to run. */
    int fundex = lookup(tokens_get_token(tokens, 0));

    if (fundex >= 0) {
      cmd_table[fundex].fun(tokens);
    } else {


    	
      /* REPLACE this to run commands as programs. */
    	int wstatus;
	    pid_t pid = fork();
      
	    if(pid){
	    	wait(&wstatus);
	        
	    }else{

		    executePrograms(line);
	    }
    }

    if (shell_is_interactive)
      /* Please only print shell prompts when standard input is not a tty */
      fprintf(stdout, "%d: ", ++line_num);

    /* Clean up memory */
    tokens_destroy(tokens);


    // setpgid(getpid(), getpid());
    tcsetpgrp(0, getpid());
    signal(SIGINT, SIG_IGN);
    signal(SIGQUIT, SIG_IGN);
    signal(SIGTSTP, SIG_IGN);
    signal(SIGCONT, SIG_IGN);
    signal(SIGTTIN, SIG_IGN);
    signal(SIGTTOU, SIG_IGN);

     printf("shell pid: %d, shell pgid: %d shell foreground pgid: %d\n", getpid(), getpgid(getpid()), tcgetpgrp(0));


  }

  return 0;
}
