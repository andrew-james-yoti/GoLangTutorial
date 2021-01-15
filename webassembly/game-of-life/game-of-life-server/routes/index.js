var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Game Of Life', subTitle: 'Golang Wasm - Game Of Life' });
});

module.exports = router;
