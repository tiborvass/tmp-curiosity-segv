#include <stdio.h>
#include <signal.h>
#include <stdlib.h>
#include <unistd.h>
#include <math.h>

static void sighandler(int signum) {
    fprintf(stderr, "Signal caught, exiting!\n");
    fprintf(stderr, "acos = %lf\n", acos(3.14));
    fprintf(stderr, "Past Go call!\n");
    exit(0);
}

int main() {
    struct sigaction sigact;
    sigact.sa_handler = sighandler;
    sigemptyset(&sigact.sa_mask);
    sigact.sa_flags = 0;
    sigaction(SIGINT, &sigact, NULL);
    sigaction(SIGUSR1, &sigact, NULL);
    sigaction(SIGTERM, &sigact, NULL);
    sigaction(SIGQUIT, &sigact, NULL);
    sigaction(SIGPIPE, &sigact, NULL);

    fprintf(stderr, "Loaded; waiting for SIGINT\n");
    acos(3.14);
    asin(3.14);

    sleep(0xFFFFFFFF);
}
