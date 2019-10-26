# Web services performance landmark values
Values for comparing your service with de-facto standard services.  I have got these values at 2xCore and 4Gb RAM VPS. 

My decision is 40 000 RPS enough value for simple web microservices.  Of course, the value depends on the logic part of your programme also. 


| RPS   |    8b  |   1k  |  10K  |  100K |  1M  | 
|-------|--------|-------|-------|-------|------| 
| nginx | ~65000 |~43000 |~41000 |~16000 |~2700 |
|   go  | ~49000 |-------|-------|-------|------|
|go-fast| ~90000 |-------|-------|-------|------|
|python3|  ~1200 |-------|-------|-------|------|
|py-fast| ~68000 |-------|-------|-------|------|
|rust   |  ~3000 |-------|-------|-------|------|


For every service, you could see how it was measured. 
