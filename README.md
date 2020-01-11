# Tadpoles.com Image Backup

#### **This is still a work in progress! - Non-functional**

Inspired by https://github.com/twneale/tadpoles but reworked from scratch to make use of the REST API behind the tadpoles.com website. 

## About
This tool will allow you to save all your child's images at full resolution from tadpoles.com.  It can be be configured with multiple save backends or new ones can be written.

Current save backends:
* Local file system
* Amazon S3 (non-functional)
* Backblaze B2 (non-functional)

## Usage
_It is preferred to run this application inside a python virtual env._

You need an authentication token. This can be easily obtained by logging into `tadpoles.com` and inspecting the cookies in any request (after login). Tokens seem to expire after about 1 month.

The requests will include a value like this:
```
Cookie: DgU00="<token_string>"
```

1. Create a file named `.env` in the root of the project.
2. Add an env variable with the contents of the auth token cookie variable without any quotes.
```
# .env file
OAUTH_TOKEN=<some_long_token_string>"

# Example:
# If in request cookie, DgU00="ABCDE"
# put this in .env file
OAUTH_TOKEN=ABCDE
```
Execute the run.py script.

## Notes
The tool queries one months worth of pictures at a time. If a query comes up with no data, then execution will stop. This may be unexpected in cases of parents having a gap of greater than 1 month in their child's attendance. This check can be skipped with the `.env` file by adding a `SKIP_NO_DATA_CHECK=true` line.

The tool will only query back for the past 10 years from the current date. This value can be configured with the `.env` file by adding a `MAX_YEARS=<int>` line.