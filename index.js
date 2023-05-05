require('dotenv').config()

const express = require('express');
const app = express();
const morgan = require('morgan')
const bodyParser = require('body-parser');

const port = process.env.PORT || 4444;

// import routes

app.use(morgan('dev'));
app.use(bodyParser.urlencoded(extended: false));
app.use(bodyParser.json());

app.use((req, res, next) => {
	res.header('Access-Control-Allow-Origin', '*');
	res.header('Access-Control-Allow-Headers', '*');

	if (req.method === 'OPTIONS') {
		res.header('Access-Control-Allow-Methods', 'POST, DELETE, GET, PUT');
		return res.status(200).json({});
	}
	next();
});

// map routes

app.use((req, res, next) => {
	const error = new Error('Notfound');
	error.status = 404;
	next(error);
});

app.use((error, req, res, next) => {
	res.status(error.status || 500);
	res.json({
		"error": {
			message: error.message
		}
	});
});

app.listen(port, () => {
	console.log("listening on port ", port)
})
