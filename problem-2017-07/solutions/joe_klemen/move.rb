module Mario
  class Move
    attr_accessor :from, :to, :coins

    def initialize(from, to = nil, coins = {})
      self.from = from
      self.to = to
      self.coins = coins
    end

    def moved?
      from != to
    end

    def inspect
      "#<Move from=#{from} to=#{to} coins=#{coins}"
    end

    def to_s
      inspect
    end
  end
end