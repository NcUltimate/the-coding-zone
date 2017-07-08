require 'active_support/all'

class Result
  attr_accessor :x, :y, :coins

  def initialize(x = 0, y = 0, coins = 0)
    @x = x
    @y = y
    @coins = coins
  end

  def location
    [x, y]
  end
end

class MiniMario
  attr_accessor :movements, :scroll, :map, :size, :start

  BLOCK_CODE = 'B'.freeze
  COIN_CODE  = 'C'.freeze

  def initialize(file_path)
    @map = []

    ingest_map!(file_path)

    @scroll = start[0] - 1
    @movements = {}

    first_fall
    puts move!
  end

  private

  def line_to_map_coords(line)
    line_data = line.split(' ')
    x = line_data[0].to_i
    y = line_data[1].to_i
    map[x] ||= []
    map[x][y] = line_data[2]
  end

  def ingest_map!(file_path)
    begin
      f = File.open(file_path, 'r')
    rescue
      raise 'Could not open file.'
    end

    @size =  f.readline.split(' ').map(&:to_i)
    @start = f.readline.split(' ').map(&:to_i)

    f.each_line { |l| line_to_map_coords(l) }
  end

  def first_fall
    add_results_to_movements([fall_and_jump(start[0], start[1])])
  end

  def max_coins_at_scroll(x)
    max_coins = 0
    movements[x - 1].each do |location, coins|
      max_coins = [max_coins, coins].max
    end
    max_coins
  end

  def add_results_to_movements(results)
    results.each do |result|
      l = result.location
      x = result.x
      movements[x] ||= {}
      movements[x][l] = [movements[x][l], result.coins].map(&:to_i).max
    end
  end

  def move!
    loop do
      @scroll += 1
      return max_coins_at_scroll(scroll) if scroll == size[0]
      break if movements[scroll]
    end

    results = {}
    movements[scroll].each do |location, coins|
      x = location[0]
      y = location[1]

      # walk
      w = walk(x, y)
      unless w.location == location
        results[w.location] = Result.new(
          w.x,
          w.y,
          [w.coins + coins, results[w.location].try(:coins)].map(&:to_i).max
        )
      end
        
      # high jump
      hj = high_jump(x, y)
      unless hj.location == location
        results[hj.location] = Result.new(
          hj.x,
          hj.y,
          [hj.coins + coins, results[hj.location].try(:coins)].map(&:to_i).max
        )
      end

      # long jump
      lj = long_jump(x, y)
      unless lj.location == location
        results[lj.location] = Result.new(
          lj.x,
          lj.y,
          [lj.coins + coins, results[lj.location].try(:coins)].map(&:to_i).max
        )
      end
    end

    add_results_to_movements(results.values)
    move!
  end

  def is_coin?(x, y)
    return false unless map[x] && map[x][y]
    return true if map[x][y] == COIN_CODE
  end

  def is_block?(x, y)
    return true if y < 0
    return true if x >= size[0] || y >= size[1]
    return false unless map[x] && map[x][y]
    return true if map[x][y] == BLOCK_CODE
  end

  def can_move_horizontally?(x, y)
    return true unless is_block?(x + 1, y)
  end
    
  def jump(x, y)
    result = Result.new(x, y)

    result.coins += 1 if is_coin?(x, y + 1)
    result.coins += 1 if !is_block?(x, y + 1) && is_coin?(x, y + 2)

    result
  end

  def fall_and_jump(x, y)
    result = Result.new(x, y)

    # proccess fall
    initial_y = y
    while !is_block?(x, y - 1)
      y -= 1
      result.coins += 1 if is_coin?(x, y)
    end
    result.y = y

    # jump if delta y is 0 or -1
    result.coins += jump(x, y).coins if initial_y == y
    result.coins += 1 if initial_y - y == 1 && is_coin?(x, y + 2)

    result
  end

  def walk(x, y)
    result = Result.new(x, y)
    return result unless can_move_horizontally?(x, y)

    # walk forward
    x += 1
    result.coins += 1 if is_coin?(x, y)

    # process potential fall
    fall_result = fall_and_jump(x, y)
    fall_result.coins += result.coins
    fall_result
  end

  def high_jump(x, y)
    result = Result.new(x, y)
    return result if is_block?(x, y + 1) || is_block?(x, y + 2)

    x = x + 1
    y = y + 3
    return result if is_block?(x, y)
    result.coins += 1 if is_coin?(x, y)

    # process fall
    fall_result = fall_and_jump(x, y)
    fall_result.coins += result.coins
    fall_result
  end

  def long_jump(x, y)
    result = Result.new(x, y)

    initial_y = y
    x += 1
    y += 1
    return result if is_block?(x, y)
    result.coins += 1 if is_coin?(x, y)

    2.times do
      break unless can_move_horizontally?(x, y)
      x += 1
      result.coins += 1 if is_coin?(x, y)
    end

    # process fall
    fall_result = fall_and_jump(x, y)
    fall_result.coins += result.coins
    fall_result
  end
end

MiniMario.new('input.txt')
