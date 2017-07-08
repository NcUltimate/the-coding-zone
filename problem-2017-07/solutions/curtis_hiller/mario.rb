class MiniMario
  attr_accessor :movements, :scroll, :map, :size, :start, :file_path

  BLOCK_CODE = 'B'.freeze
  COIN_CODE  = 'C'.freeze

  def initialize(file_path)
    @file_path = file_path
    @map = []

    ingest_map!

    @scroll = start[0] - 1
    @movements = { start[0] => { start => 0 } }

    first_jump
    puts move!
  end

  private

  def first_jump
    if is_block?(start[0], start[1] - 1)
      result = jump(start[0], start[1])
    else
      result = fall(start[0], start[1])
      
      # fall if delta is only -1
      if start[1] == result[:location][1] + 1
        result[:coins] += 1 if is_coin?(start[0], start[1] + 1)
      end
    end
    
    add_results_to_movements([[result[:location], result[:coins]]])
  end

  def move!
    loop do
      @scroll += 1
      return max_coins_at_scroll(scroll) if scroll == size[0]
      break if movements[scroll]
    end

    movements[scroll].each do |location, coins|
      results = {}
      x = location[0]
      y = location[1]

      # walk and jump
      waj_results = walk_and_jump(x, y)
      unless waj_results[:location] == location
        results[waj_results[:location]] = waj_results[:coins] + coins
      end
        
      # high jump
      hj_results = high_jump(x, y)
      l = hj_results[:location]
      unless l == location
        results[l] = [results[l], hj_results[:coins] + coins].map(&:to_i).max
      end

      # long jump
      lj_results = long_jump(x, y)
      l = lj_results[:location]
      unless l == location
        results[l] = [results[l], lj_results[:coins] + coins].map(&:to_i).max
      end

      add_results_to_movements(results)
    end

    move!
  end

  def add_results_to_movements(results)
    results.each do |result|
      l = result[0]
      x = l[0]
      movements[x] ||= {}
      movements[x][l] = [movements[x][l], result[1]].map(&:to_i).max
    end
  end

  def ingest_map!
    begin
      f = File.open(file_path, 'r')

      @size =  f.readline.split(' ').map(&:to_i)
      @start = f.readline.split(' ').map(&:to_i)

      f.each_line do |line|
        line_data = line.split(' ')
        x = line_data[0].to_i
        y = line_data[1].to_i
        map[x] ||= []
        map[x][y] = line_data[2]
      end
    rescue
      raise 'Could not open file.'
    end
  end

  def max_coins_at_scroll(x)
    x -=  1

    max_coins = 0
    movements[x].each do |location, coins|
      max_coins = [max_coins, coins].max
    end
    max_coins
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

  def fall(x, y)
    results = { location: [x, y], coins: 0 }

    while(y > 0)
      y -= 1
      break if is_block?(x, y)
      results[:location] = [x, y]
      results[:coins] += 1 if is_coin?(x, y)
    end

    results
  end

  def walk_and_jump(x, y)
    return { location: [x, y], coins: 0 } unless can_move_horizontally?(x, y)

    results = walk(x, y)
    x += 1

    # jump if those square haven't been covered to avoid duplicate coins
    if results[:location][1] == y
      jump_results = jump(x, y)
      results[:coins] += jump_results[:coins]
      return results
    end
    if results[:location][1] == y - 1
      results[:coins] += 1 if is_coin?(x, y + 2)
    end
    results
  end

  def walk(x, y)
    results = { location: [x, y], coins: 0 }
    return results unless can_move_horizontally?(x, y)

    # walk forward
    x += 1
    results[:coins] += 1 if is_coin?(x, y)

    # process potential fall
    fall_results = fall(x, y)

    {
      location: fall_results[:location],
      coins: results[:coins] + fall_results[:coins]
    }
  end
    
  def jump(x, y)
    results = { location: [x, y], coins: 0 }

    results[:coins] += 1 if is_coin?(x, y + 1)
    if !is_block?(x, y + 1) && is_coin?(x, y + 2)
      results[:coins] += 1
    end

    results
  end

  def high_jump(x, y)
    results = { location: [x, y], coins: 0 }
    return results if is_block?(x, y + 1) || is_block?(x, y + 2)

    x = x + 1
    y = y + 3
    return results if is_block?(x, y)
    new_y = y
    results[:coins] += 1 if is_coin?(x, y)

    # process fall
    fall_results = fall(x, y)
    x = fall_results[:location][0]
    y = fall_results[:location][1]

    # jump required if delta y is 0 or -1
    if y == new_y - 1
      fall_results[:coins] += 1 if is_coin?(x, y + 2)
    elsif y == new_y
      jump_results = jump(x, y)
      fall_results[:coins] += jump_results[:coins]
    end

    {
      location: [x, y],
      coins: results[:coins] + fall_results[:coins]
    }
  end

  def long_jump(x, y)
    results = { location: [x, y], coins: 0 }

    initial_y = y
    x += 1
    y += 1
    return results if is_block?(x, y)
    results[:coins] += 1 if is_coin?(x, y)

    2.times do
      break if !can_move_horizontally?(x, y)
      x += 1
      results[:coins] += 1 if is_coin?(x, y)
    end

    # process fall
    fall_results = fall(x, y)
    x = fall_results[:location][0]
    y = fall_results[:location][1]
    
    # jump required if delta y is 0 or +1
    if y == initial_y
      fall_results[:coins] += 1 if is_coin?(x, y + 2)
    elsif y > initial_y
      jump_results = jump(x, y)
      fall_results[:coins] += jump_results[:coins]
    end

    {
      location: [x, y],
      coins: results[:coins] + fall_results[:coins]
    }
  end
end

mario = MiniMario.new('input.txt')
