var reader = require('fs');

(function () {
  var size;
  var lines = [];
  var words = [];
  var highlights = [];
  var cardinals = {
    right:     '',
    left:      '',
    up:        '',
    down:      ''
  }
  var diagonals = {
    upleft:    [],
    downleft:  [],
    upright:   [],
    downright: []
  }

  function output(){
    var rows = highlights.length;
    var cols;
    var i = 0;
    var j;
    var output;

    for (; i < rows; i++){
      output = '';
      j = 0;
      for (; j < highlights[i].length; j++){
        output += highlights[i][j];
      }
      console.log(output);
    }
  }

  function highlightLetter(y, x){
    newLetter  = '\x1b[44m\x1b[30m';
    newLetter += lines[y].charAt(x);
    newLetter += '\x1b[0m\x1b[0m';
    highlights[y][x] = newLetter;
  }

  function findWordLetters (direction, idx, word, array_node){
    var length = word.length;
    var x, y;
    var delta_x = delta_y = 0;

    switch (direction){
      case 'right':
        delta_x = 1;
        x = idx % size[1];
        y = Math.floor(idx / size[1]);
      break;

      case 'left':
        delta_x = -1;
        y = (size[0] - 1) - Math.floor(idx / size[0]);
        x = (size[1] - 1) - (idx % size[1]);
      break;

      case 'up':
        // console.log('so...', idx);
        delta_y = -1;
        x = Math.floor(idx / size[0]);
        y = (size[0] - 1) - (idx % size[0]);
      break;

      case 'down':
        delta_y = 1;
        x = Math.floor(idx / size[0]);
        y = idx % size[0];
      break;

      case 'upright':
        delta_y = -1;
        delta_x = 1;
        //-lower half
        if (array_node < size[0]){
          y = array_node;
          x = idx;
        }
        //-upper half
        else{
          y = size[0] - 1 - idx;
          x = (array_node - (size[0] - 1)) + idx;
        }
      break;

      case 'downright':
        delta_y = 1;
        delta_x = 1;
        //-lower half
        if (array_node < size[0]){
          y = size[0] - 1 - array_node;
          x = idx;
        }
        //-upper half
        else{
          y = idx;
          x = (array_node - (size[0] - 1)) + idx;
        }
      break;

      case 'upleft':
        delta_y = -1;
        delta_x = -1;
        //-lower half
        if (array_node < size[0]){
          y = array_node;
          x = size[1] - 1 - idx;
        }
        //-upper half
        else{
          y = size[0] - 1 - idx;
          x = (size[1] - 2) - idx - (array_node - (size[1] - 1));
        }
      break;

      //-downleft
      default:
        delta_y = 1;
        delta_x = -1;
        //-lower half
        if (array_node < size[0]){
          y = size[0] - 1 - array_node + idx;
          x = size[1] - 1 - idx;
        }
        else{
          y = idx;
          x = (size[1] - 2) - idx - (array_node - (size[1] - 1));
        }
      break;
    }

    // console.log('word was found at (', y, ',', x, ')');

    var i = 0;
    for(; i < length; i++){
      highlightLetter(y + (delta_y * i), x + (delta_x * i));
    }
  }

  function findWord(word){
    var idx, s, i, run;
    // console.log('finding ', word);

    //-try cardinals
    for (s in cardinals){
      // console.log('finding word in', cardinals[s]);
      idx = cardinals[s].indexOf(word);
      if (idx !== -1){
        // console.log('word was found in', s, 'at', idx);
        return findWordLetters(s, idx, word);
      }
    }

    //-try diagonals (only array members long enough)
    for (s in diagonals){
      i = word.length - 1;
      run = diagonals[s].length - i;
      // console.log('trying ', s);
      for (; i < run;){
        // console.log('comparing', word, 'with', diagonals[s][i]);
        idx = diagonals[s][i].indexOf(word);
        if (idx !== -1){
          // console.log('word was found at', idx, 'on row', i);
          return findWordLetters(s, idx, word, i);
        }
        i++;
      }
    }
    return false;
  }

  function solve (){
    var run = words.length;
    var i = 0;
    for (; i < run; i++){
      findWord(words[i]);
    }
  }

  function handleFileLoad (err, data){
    var x, y, y2, y_mark, y_mark_2, i;

    lines = data.split('\n');

    size = lines.shift().split(' ');
    size[0] = parseInt(size[0]);
    size[1] = parseInt(size[1]);

    while (lines.length > size[0]){
      words.push(lines.pop());
    }

    y = 0;
    for (; y < lines.length; y++){
      highlights[y] = lines[y].split('');
    }

    cardinals.right = lines.join('');
    cardinals.left = cardinals.right.split('').reverse().join('');

    //-up and down cardinals
    x = 0;
    for (; x < size[1]; x++){
      y = 0;
      y2 = size[0] - 1;
      for (; y < size[0]; y++, y2--){
        cardinals.down += lines[y].charAt(x);
        cardinals.up += lines[y2].charAt(x);
      }
    }

    //-diagonals starting from upper left and lower left
    y_mark = 0;
    y_mark_2 = size[0] - 1;
    while (y_mark < size[0]){
      y = y_mark;
      y2 = y_mark_2;
      x = 0;
      for (; y >= 0;){
        diagonals.upright += lines[y].charAt(x);
        diagonals.downright += lines[y2].charAt(x);
        y--;
        y2++;
        x++;
        if (x > size[1]){
          break;
        }
      }
      diagonals.upright += ' ';
      diagonals.downright += ' ';
      y_mark++;
      y_mark_2--;
    }
    //-starting from lower left +1 x and upper left +1 x
    x_mark = 1;
    while (x_mark < size[1]){
      y = size[0] - 1;
      y2 = 0;
      x = x_mark;
      for (; x < size[1];){
        diagonals.upright += lines[y].charAt(x);
        diagonals.downright += lines[y2].charAt(x);
        y2++;
        y--;
        x++;
        if (y < 0){
          break;
        }
      }
      diagonals.upright += ' ';
      diagonals.downright += ' ';
      x_mark++;
    }

    diagonals.downleft = diagonals.upright.trim().split('').reverse().join('');
    diagonals.upleft = diagonals.downright.trim().split('').reverse().join('');
    diagonals.upright = diagonals.upright.trim().split(' ');
    diagonals.downright = diagonals.downright.trim().split(' ');
    diagonals.upleft = diagonals.upleft.split(' ');
    diagonals.downleft = diagonals.downleft.split(' ');

    solve();
    output();
  }

  function initialize (fileName){
    reader.onload = handleFileLoad;
    reader.readFile(fileName, 'utf8', handleFileLoad);
  }

  initialize('input.txt');
}());
