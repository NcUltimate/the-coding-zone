require 'active_support/all'

def test(pieces)
  raise 'Board not legal!' unless board_legal?

  @all_pieces = []
  @bishops = []
  @rooks = []
  @bk = []
  @wk = []

  pieces.each { |piece| log_piece(piece) }
  check_rooks!
  check_bishops!
end

private

def board_legal?
  true
end

def coords_from_pos(str)
  [str[0].downcase.ord-96, str[1].to_i]
end

def log_piece(p)
  cd = p.last(2)
  @all_pieces << coords_from_pos(cd)
  
  return @bishops << coords_from_pos(cd) if p.include?('WB')
  return @rooks << coords_from_pos(cd) if p.include?('WR')
  return @bk = coords_from_pos(cd) if p.include?('BK')
  @wk = coords_from_pos(cd)
end

def bishop_check_squares
  return @bishop_check_squares if @bishop_check_squares
  @bishop_check_squares = []

  # lower left to upper right
  if @bk[0] < @bk[1]
    (1..7).each do |i|
      y = @bk[1] - @bk[0] + i
      break if y > 8
      @bishop_check_squares << [i, y] unless @bk == [i, y]
    end
  else
    (1..7).each do |i|
      x = @bk[0] - @bk[1] + i
      break if x > 8
      @bishop_check_squares << [x, i] unless @bk == [i, y]
    end
  end

  # lower right to upper left
  (1..8).reverse_each do |i|
    y = @bk[0] - (i - 7)
    break if y > 8
    @bishop_check_squares << [i, y] unless @bk == [i, y]
  end

  @bishop_check_squares
end

def rook_check_squares
  return @rook_check_squares if @rook_check_squares
  @rook_check_squares = []

  (1..8).each do |i|
    y = @bk[1]
    x = @bk[0]
    @rook_check_squares << [i, y] unless @bk == [i, y]
    @rook_check_squares << [x, i] unless @bk == [i, y]
  end

  @rook_check_squares
end

def king_moves(king)
  moves = []
  x = king[0]
  y = king[1]
  i = x + 1
  j = y + 1
  moves << [i, y] unless i > 8 # 2, 1
  moves << [i, j] unless i > 8 || j > 8 # 2, 2
  moves << [x, j] unless j > 8 # 1, 2
  i = x - 1
  moves << [i, j] unless i < 1 || j > 8 # 0, 2
  j = y - 1
  moves << [i, y] unless i < 1 # 0, 1
  moves << [i, j] unless i < 1 || j < 1 # 0, 0
  moves << [x, j] unless j < 1 # 1, 0
  i = x + 1
  moves << [i, j] unless i > 8 || j < 1 # 2, 0
  moves
end

def black_king_moves
  return @black_king_moves if @black_king_moves

  @black_king_moves = king_moves(@bk)
end

def white_king_moves
  return @white_king_moves if @white_king_moves

  @white_king_moves = king_moves(@wk)
end

def mate!(move)
  puts 'mate!'
  puts move
end

def square_threatened?(square, rooks, bishops)
  new_pieces = rooks + bishops + [@wk] + [@bk]
  rooks.each do |rook|
    return true if rook_moves(rook, new_pieces).include?(square)
  end
  bishops.each do |bishop|
    return true if bishop_moves(bishop, new_pieces).include?(square)
  end
  return true if white_king_moves.include?(square)
  false
end

def king_is_mated?(rooks, bishops)
  black_king_moves.each do |move|
    return false unless square_threatened?(move, rooks, bishops)
  end
  true
end

def try_rook_mates!(rook, checks)
  other_rooks = @rooks - [rook]
  checks.each do |check|
    mate!("R -> #{check[0]}, #{check[1]}") if king_is_mated?(other_rooks << check, @bishops)
  end
end

def rook_moves(rook, pieces)
  x = rook[0]
  y = rook[1]
  moves = []
  i = x + 1
  while i < 9 && !pieces.include?([i, y])
    moves << [i, y]
    i += 1
  end
  i = x - 1
  while i > 0 && !pieces.include?([i, y])
    moves << [i, y]
    i -= 1
  end
  i = y + 1
  while i < 9 && !pieces.include?([x, i])
    moves << [x, i]
    i += 1
  end
  i = y - 1
  while i > 0 && !pieces.include?([x, i])
    moves << [x, i]
    i -= 1
  end
  moves
end

def rook_checks(rook)
  return [] unless rook_check_squares
  check_moves = rook_check_squares & rook_moves(rook, @all_pieces)
  return [] unless check_moves
  x = @bk[0]
  y = @bk[1]
  check_moves.each do |move|
    delete = false
    if move[0] == x
      if move[1] < y
        (move[1]+1...y).each do |j|
          delete = true if @all_pieces.include?([x, j])
        end
      else
        (y+1...move[1]).each do |j|
          delete = true if @all_pieces.include?([x, j])
        end
      end
    else
      if move[0] < x
        (move[0]+1...x).each do |i|
          delete = true if @all_pieces.include?([i, y])
        end
      else
        (x+1...move[0]).each do |i|
          delete = true if @all_pieces.include?([i, y])
        end
      end
    end
    check_moves.delete(move) if delete
  end
  check_moves
end

def check_rooks!
  return unless @rooks.any?
  @rooks.each do |rook|
    checks = rook_checks(rook)
    next unless checks.any?
    try_rook_mates!(rook, checks)
  end
end

def try_bishop_mates!(bishop, checks)
  other_bishops = @bishops - [bishop]
  checks.each do |check|
    mate!("B -> #{check[0]}, #{check[1]}") if king_is_mated?(@rooks, other_bishops << check)
  end
end

def bishop_moves(bishop, pieces)
  x = bishop[0]
  y = bishop[1]
  moves = []
  i = x + 1
  j = y + 1
  while i < 9 && j < 9 && !pieces.include?([i, j])
    moves << [i, j]
    i += 1
    j += 1
  end
  i = x - 1
  j = y - 1
  while i > 0 && j > 0 && !pieces.include?([i, j])
    moves << [i, j]
    i -= 1
    j -= 1
  end
  i = x + 1
  j = y - 1
  while i < 9 && j > 0 && !pieces.include?([i, j])
    moves << [i, j]
    i += 1
    j -= 1
  end
  i = x - 1
  j = y + 1
  while i > 0 && j < 9 && !pieces.include?([i, j])
    moves << [i, j]
    i -= 1
    j += 1
  end
  moves
end

def bishop_checks(bishop)
  return [] unless @bishop_check_squares
  check_moves = @bishop_check_squares & bishop_moves(bishop, @all_pieces)
  return [] unless check_moves
  x = @bk[0]
  y = @bk[1]
  check_moves.each do |move|
    delete = false
    if move[0] > x && move[1] > y
      i = move[0] - 1
      j = move[1] - 1
      while i > x && j > y
        delete = true if @all_pieces.include?([i, j])
        i -= 1
        j -= 1
      end
    end
    if move[0] < x && move[1] > y
      i = move[0] + 1
      j = move[1] - 1
      while i < x && j > y
        delete = true if @all_pieces.include?([i, j])
        i += 1
        j -= 1
      end
    end
    if move[0] > x && move[1] < y
      i = move[0] - 1
      j = move[1] + 1
      while i > x && j < y
        delete = true if @all_pieces.include?([i, j])
        i -= 1
        j += 1
      end
    end
    if move[0] < x && move[1] < y
      i = move[0] + 1
      j = move[1] + 1
      while i < x && j < y
        delete = true if @all_pieces.include?([i, j])
        i += 1
        j += 1
      end
    end
    check_moves.delete(move) if delete
  end
  check_moves
end

def check_bishops!
  return unless @bishops.any?
  @bishops.each do |bishop|
    checks = bishop_checks(bishop)
    try_bishop_mates!(bishop, checks)
  end
  false
end


test([
  'BK a1',
  'WK h8',
  'WB c4',
  'WR b3',
  'WB d3',
  'WR b8'
])
