module Mario
  class Algorithm
    class << self
      def run!(grid)
        self.new(grid).run!
      end
    end

    attr_accessor :grid
    def initialize(grid)
      self.grid = grid
    end

    # Dynamic Programming
    def run!
      # Initialize DP variables. We are solving the subproblem
      # of "from a given space (x, y), what is the maximum number
      # of coins between this space and the end of the level"
      max_coins = {}
      collected_by = {}

      # In reverse order, Iterate over all spaces
      # that are directly over blocks
      spaces = grid.walkable_spaces.sort_by(&:x)
      spaces.reverse.each do |space|
        collected_by[space.hash] ||= []

        # Perform all 3 forward-moving moves.
        moves = [
          std_move_from(space),
          high_jump_from(space),
          long_jump_from(space)
        ]

        # Find which of these moves collects the most coins
        # that haven't been "collected_by" any move yet.
        best_move = moves.max_by do |end_pos, coins|
          next -1 if end_pos == space
          uncollected = coins - collected_by[end_pos.hash]
          max_coins[end_pos.hash].to_i + uncollected.size
        end
        
        uncollected = best_move[1] - collected_by[best_move[0].hash]

        # Update the largest number of coins for the given space
        # to be coins collected plus the max of the space we land in.
        max_coins[space.hash] =
          max_coins[best_move[0].hash].to_i + uncollected.size

        # Remember that we collected the coins
        # for subsequent iterations.
        best_move[1].each do |coin|
          collected_by[space.hash] << coin
        end
      end

      # Make our player fall until he is in
      # a space above a block. It follows that he MUST land in a
      # space that we computed the max value for in the above loop.
      #
      # Collect coins along the way. The answer to the entire grid
      # is max_coins[start].
      start, coins = fall_and_collect(grid.start)
      max_coins[start.hash] + (coins - collected_by[start.hash]).size
    end

    private

    def std_move_from(p)
      jump_then_move(p) do
        path = [Point[1, 0]]
        fall_and_collect(*trace_and_collect(p, *path))
      end
    end

    def high_jump_from(p)
      jump_then_move(p) do
        path = [Point[0, 1], Point[0, 1], Point[1, 1]]
        fall_and_collect(*trace_and_collect(p, *path))
      end
    end

    def long_jump_from(p)
      jump_then_move(p) do
        path = [Point[1, 1], Point[1, 0], Point[1, 0]]
        fall_and_collect(*trace_and_collect(p, *path))
      end
    end

    def std_jump_from(p)
      path = [Point[0, 1], Point[0, 1]]
      fall_and_collect(*trace_and_collect(p, *path))
    end

    def jump_then_move(p)
      t1, c1 = std_jump_from(p)
      t2, c2 = yield
      [t2, c1 | c2]
    end

    def fall_and_collect(p, c = [])
      end_pos, coins = trace_and_collect(p, Point[0, -1])
      end_pos == p ? [p, c + coins] : fall_and_collect(end_pos, c + coins)
    end

    def trace_and_collect(p, *steps)
      end_pos = p
      coins = []

      steps.each do |step|
        prev = end_pos
        end_pos = prev + step
        return [prev, coins] unless grid.include?(end_pos)
        return [prev, coins] if grid.block_at?(end_pos)
        coins << end_pos if grid.coin_at?(end_pos)
      end
      [end_pos, coins]
    end
  end
end