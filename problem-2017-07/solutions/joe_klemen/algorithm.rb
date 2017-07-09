module Mario
  class Algorithm
    def self.run!(grid)
      self.new(grid).run!
    end

    attr_accessor :grid, :max_coins, :collected_at
    def initialize(grid)
      self.grid = grid
      self.max_coins = {}
      self.collected_at = {}
    end

    # Using Dynamic Programming, we are solving the subproblem
    # of "from a given space (x, y), what is the maximum number
    # of coins between this space and the end of the level?"
    def run!
      # In reverse order, Iterate over all spaces
      # that are directly over blocks
      spaces = grid.walkable_spaces.sort_by(&:x)
      spaces.reverse.each do |space|
        # Get the best move
        best_move = best_move_from(space)

        # Update the largest number of coins for the given space
        # to be coins collected plus the max of the space we land in.
        max_coins[best_move.from.hash] =
          max_coins[best_move.to.hash].to_i + num_uncollected(best_move)

        # Remember where we collected each coin
        # for subsequent iterations.
        best_move.coins.each do |coin_hash, _val|
          collected_at[coin_hash] ||= best_move.from.x
        end
      end

      # Make our player fall until he is in a space above a block. 
      # Collect coins along the way. The answer to the entire grid
      # is max_coins from where he falls to plus whatever he collects
      # along the way that has not already been collected.
      fall = PathFinder.fall(grid, grid.start)
      max_coins[fall.to.hash] + num_uncollected(fall)
    end

    private

    # Find the move from a given space that
    # collects the most coins
    def best_move_from(space)
      moves_from(space).max_by do |move|
        next -1 unless move.moved?
        max_coins[move.to.hash].to_i + num_uncollected(move)
      end
    end

    # Perform all 3 forward-moving moves.
    def moves_from(space)
      [
        PathFinder.std_move(grid, space),
        PathFinder.high_jump(grid, space),
        PathFinder.long_jump(grid, space)
      ]
    end

    # For a given move, how many of its coins
    # have not been collected already
    def num_uncollected(move)
      move.coins.count do |coin_hash, _val|
        !collected_at[coin_hash] || collected_at[coin_hash] != move.to.x
      end
    end
  end
end