


```
uvicorn main:app --reload
gunicorn wait$:app --workers 4 --worker-class uvicorn.workers.UvicornWorker --bind 0.0.0.0:8000
gunicorn main:app --workers 4 --worker-class uvicorn.workers.UvicornWorker --bind 0.0.0.0:8000
```
