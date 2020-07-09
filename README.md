# Counter
A request counter service base on echo framework.

The ipList in EchoEngine is used for clustering.

The [getRequest](https://morning-taiga-80604.herokuapp.com/) will get the number of requests for this IP in one minute.

The [getAllRequest](https://morning-taiga-80604.herokuapp.com/all) will get all the number of requests for all IPs in one minute.

An error will occur if the limit is exceeded.