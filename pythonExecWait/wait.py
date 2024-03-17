from fastapi import FastAPI
from random import randint
import time

app = FastAPI()


@app.get("/")
async def root():
    start = time.ctime()
    print(f"Receiving request at {start}")
    sleeptime = randint(1,10)
    time.sleep(sleeptime)
    end = time.ctime()
    return {"message": f"Query received at {start} / waitted {sleeptime} sec / response at {end}  "}