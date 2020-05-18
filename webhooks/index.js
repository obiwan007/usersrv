var express = require("express");
var app = express();
var childProcess = require("child_process");
var githubUsername = "obiwan007";

app.get("/webhooks/github", function (req, res) {
  console.log("Received GET ");
});

app.post("/webhooks/github", function (req, res) {
  console.log("Request", req);
  var sender = req.body.sender;
  var branch = req.body.ref;

  if (branch.indexOf("master") > -1 && sender.login === githubUsername) {
    deploy(res);
  }
});

function deploy(res) {
  console.log("Receved trigger");
  childProcess.exec(
    "cd /home/markus/dev/usersrv && make docker && make kreapply",
    function (err, stdout, stderr) {
      if (err) {
        console.error(err);
        return res.send(500);
      }
      res.send(200);
    }
  );
}

app.listen(38080, "0.0.0.0", function () {
  console.log("Example app listening on port 38080!");
});
