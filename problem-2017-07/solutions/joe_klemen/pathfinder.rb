module Mario  
  class PathFinder 
    class << self
      def fall(grid, p)
        self.new(grid).fall(p)
      end

      def std_move(grid, p)
        self.new(grid).std_move(p)
      end

      def high_jump(grid, p)
        self.new(grid).high_jump(p)
      end

      def long_jump(grid, p)
        self.new(grid).long_jump(p)
      end
    end

    attr_reader :grid
    def initialize(grid)
      @grid = grid
    end

    def fall(p)
      path = [Point[0, -1]]
      fall_and_collect(move_and_collect(p, *path))
    end

    def std_move(p)
      path = [Point[1, 0]]
      jump_then_move(p, path)
    end

    def high_jump(p)
      path = [Point[0, 1], Point[0, 1], Point[1, 1]]
      jump_then_move(p, path)
    end

    def long_jump(p)
      path = [Point[1, 1], Point[1, 0], Point[1, 0]]
      jump_then_move(p, path)
    end

    private

    def std_jump(p)
      path = [Point[0, 1], Point[0, 1]]
      fall_and_collect(move_and_collect(p, *path))
    end

    def jump_then_move(p, path)
      jump = std_jump(p)
      fall = fall_and_collect(move_and_collect(p, *path))

      fall.coins.merge!(jump.coins)
      fall
    end

    def fall_and_collect(move)
      fall = move_and_collect(move.to, Point[0, -1])
      if fall.to == move.to
        move
      else
        move = Move.new(move.from, fall.to, move.coins.merge(fall.coins))
        fall_and_collect(move)
      end
    end

    def move_and_collect(p, *path)
      end_pos = p
      move = Move.new(p, end_pos)
      path.each do |step|
        end_pos = move.to + step
        return move unless grid.can_enter?(end_pos)

        move.coins[end_pos.hash] = true if grid.coin_at?(end_pos)
        move.to = end_pos
      end
      move
    end
  end
end