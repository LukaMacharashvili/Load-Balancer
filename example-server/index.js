const express = require("express");
const cors = require("cors");

const app = express();
const PORT = process.env.PORT || 3000;

app.use(
    cors({
        origin: "*",
    })
);

app.use(express.json({ limit: "25mb" }));

app.route("/test").get((req, res) => {
    res.send("Hello World!");
});

app.route("/health").get((req, res) => {
    res.send("OK");
});

app.listen(PORT, () => console.log("Server Started!"));