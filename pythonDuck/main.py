from fastapi import FastAPI, Request
from random import randint
import time

import duckdb


app = FastAPI()




@app.post("/exec")
async def exec(request: Request):

    query = await request.body()
    print(f"Received this query : {query}")

    con = duckdb.connect(":memory:")

    res = con.execute(query).fetchall()

    """start = time.ctime()
    print(f"Receiving request at {start}")
    sleeptime = randint(1,10)
    time.sleep(sleeptime)
    end = time.ctime()

    return {"message": f"Query received at {start} / waitted {sleeptime} sec / response at {end}  "}"""

    return {"results" : res }