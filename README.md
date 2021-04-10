# tmp-curiosity-segv

I don't know why (similar) code to this is segfaulting Go. I'm checking this in
to show friends what (basically) is going on. This fails in a different way,
but I'm guessing is the same root cause.

## Crash

```
$ make
$ LD_LIBRARY_PATH=. ./main
```

When C-c'd:

```
Loaded; waiting for SIGINT
2021/04/09 16:17:42 asin called
^CSignal caught, exiting!
fatal: morestack on g0
Trace/breakpoint trap
```


## Crash (not on this program, but similar) on amd64

```
Program terminated with signal SIGSEGV, Segmentation fault.
#0  runtime.cgocallback_gofunc () at /usr/lib/go-1.15/src/runtime/asm_amd64.s:781
781		MOVQ	(g_sched+gobuf_sp)(SI), DI  // prepare stack as DI

#0  runtime.cgocallback_gofunc () at /usr/lib/go-1.15/src/runtime/asm_amd64.s:781
#1  0x00007f8d5e31933e in runtime.cgocallback () at /usr/lib/go-1.15/src/runtime/asm_amd64.s:705
#2  0x00007f8d5e574247 in _cgoexp_4dbe752b3cc6_fn (a=0x7ffcc4c07b90, n=16, ctxt=0) at _cgo_gotypes.go:199
#3  0x00007f8d5e31d757 in crosscall2 () at /usr/lib/go-1.15/src/runtime/cgo/asm_amd64.s:59
```

## Crash (not on this program, but similar) on arm64

```
Program terminated with signal SIGSEGV, Segmentation fault.
#0  runtime.cgocallback () at /opt/golang/go-1.16/src/runtime/asm_arm64.s:1041
1041		MOVD	(g_sched+gobuf_sp)(g), R4 // prepare stack as R4
[Current thread is 1 (Thread 0xffff8b342120 (LWP 355301))]
```
