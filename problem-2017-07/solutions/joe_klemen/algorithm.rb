module Mario
  class Algorithm
    class << self
      def run!(grid)
        self.new(grid).run!
      end
    end

    attr_accessor :grid
    attr_reader :max_coins, :collected_by
    def initialize(grid)
      self.grid = grid

      # Initialize DP variables. We are solving the subproblem
      # of "from a given space (x, y), what is the maximum number
      # of coins between this space and the end of the level"
      @max_coins = {}
      @collected_by = {}
    end

    def run!
      # In reverse order, Iterate over all spaces
      # that are directly over blocks
      spaces = grid.walkable_spaces.sort_by(&:x)
      spaces.reverse.each do |space|
        collected_by[space.hash] ||= []

        # Perform all 3 forward-moving moves.
        moves = [
          PathFinder.std_move(grid, space),
          PathFinder.high_jump(grid, space),
          PathFinder.long_jump(grid, space)
        ]

        # Find which of these moves collects the most coins
        # that haven't been "collected_by" any move yet.
        best_move = moves.max_by do |move|
          next -1 unless move.moved?
          uncollected = move.coins - collected_by[move.to.hash]
          max_coins[move.to.hash].to_i + uncollected.size
        end
        
        uncollected = best_move.coins - collected_by[best_move.to.hash]

        # Update the largest number of coins for the given space
        # to be coins collected plus the max of the space we land in.
        max_coins[best_move.from.hash] =
          max_coins[best_move.to.hash].to_i + uncollected.size

        # Remember that we collected the coins
        # for subsequent iterations.
        collected_by[best_move.from.hash] = uncollected
      end

      # Make our player fall until he is in
      # a space above a block. It follows that he MUST land in a
      # space that we computed the max value for in the above loop.
      #
      # Collect coins along the way. The answer to the entire grid
      # is max_coins[start].
      fall = PathFinder.fall(grid, grid.start)
      max_coins[fall.to.hash] + (fall.coins - collected_by[fall.to.hash]).size
    end
  end
end