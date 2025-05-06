local key = KEYS[1]
local expectCode = ARGV[1]
local code = redis.call("get", key)
local cntKey = key..":cnt"
local cnt = tonumber(redis.call("get", cntKey))
if cnt <= 0 then
    return -1
if expectCode != code then
    redis.call("decr", cntKey, 1)
    return -2
else
    redis.call("del", key)
    redis.call("del", cntKey)
    return 0