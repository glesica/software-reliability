# Logging Assignment

Improve the logging situation in the example application from
the integration testing section. Specifically, you should attempt
to improve two things:

  1. Add a correlation ID so that log messages from the different systems
     can be tied together based on the request that triggered them.
  2. Using the database configuration as a template, modify the Dockerfiles and
     the `docker-compose.yaml` file so that they mount a volume for logs to be
     stored in, then modify the code to write log messages to this location.

