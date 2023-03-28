# http_tool

small custom http tool to generate parallel http requests

## requirements

to create the binary file run the following command:
    $> make build

to run all the unit test run the following command:
    $> make test

## run tool

to run the tool after create the binary file, run the following command:

    $> ./build/http_tool -parallel 7 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com

## dictionary

$> ./build/http_tool -parallel ${int_value} ${urls_list}

* -parallel ${int_value} is a optional argument that allows you to set the number of parallel requests. Where ${int_value} is a numeric value that must be a greater than zero, for example 7.

* ${urls_list} is a list of valid urls that the tool is going to make htpt request to.
for example:  adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
