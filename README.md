## The solution of 10th exam on 23.04.2023

While solving exam task there were no requirements below:

* putting "omitempty" or any other value on "json" tag on **ProcessingResult** field of **Post** model or any other
  * _it is done here in pkg/models/models.go at line 16_
* handling "Record not found" error in **GetOnePost** method
  * _it is done here in pkg/handlers/posts.go from 27 to 38 line_
* getting DB credentials from console flags
  * _it is done here in cmd/api/main.go from 14 to 18 line_