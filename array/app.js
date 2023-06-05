const { exec } = require("child_process");

const args = [1, 2, 3, 4];
exec(`go run main.go ${args.join(" ")}`, (err, stdout, stderr) => {
  if (err) {
    console.error(err);
    return;
  }
  console.log(stdout); // 출력 결과: 10
});
