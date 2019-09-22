# Use an official Python runtime as a parent image
FROM python:2-slim

# Set the working directory to /app
WORKDIR /scheduler

# Copy the current directory contents into the container at /app
COPY . /scheduler

# Install any needed packages specified in requirements.txt
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Run app.py when the container launches
CMD ["python", "scheduler.py"]


