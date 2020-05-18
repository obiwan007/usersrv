import { exec } from "child_process";
import crypto from "crypto";
import http from "http";

const SECRET = "THEDAMNSECRET";

http
  .createServer((req, res) => {
    req.on("data", (chunk) => {
      console.log("Receved data");
      const signature = `sha1=${crypto
        .createHmac("sha1", SECRET)
        .update(chunk)
        .digest("hex")}`;

      const isAllowed = req.headers["x-hub-signature"] === signature;

      const body = JSON.parse(chunk);

      const isMaster = body.ref === "refs/heads/master";

      if (isAllowed && isMaster) {
        // do something
        deploy(res);
      }
    });

    res.end();
  })
  .listen(38080);

function deploy(res) {
  console.log("Receved trigger");
  const cp = exec(
    "cd /home/markus/dev/usersrv && git pull && make docker && make kapply && make kredeploy",
    (err, stdout, stderr) => {
      if (err) {
        console.error(err);
        // return res.send(500);
      }
      console.log("Done");
      // res.send(200);
    }
  );
  cp.stdout.pipe(process.stdout);
  cp.stderr.pipe(process.stderr);
}
