module Intersections
  class Algorithm
    def self.run!(poly1, poly2)
      self.new(poly1, poly2).run!
    end

    attr_accessor :poly1, :poly2, :manifest
    def initialize(poly1, poly2)
      self.poly1 = poly1
      self.poly2 = poly2
      self.manifest = {}
    end

    def run!
      initialize_with_polys
      puts "==============="
      poly1.edges.each do |edge1|
        poly2.edges.each do |edge2|
          int = edge1.intersection(edge2)
          next unless int

          if edge1.split_by?(int)
            add_entry(Line.new(int, edge1.p1))
            add_entry(Line.new(int, edge1.p2))
          end

          if edge2.split_by?(int)
            add_entry(Line.new(int, edge2.p1))
            add_entry(Line.new(int, edge2.p2))
          end
        end
      end
      manifest
    end

    private

    def initialize_with_polys
      poly1.edges.each(&method(:add_entry))
      poly2.edges.each(&method(:add_entry))
    end

    def add_entry(edge)
      puts "Adding #{edge} #{edge.p1.hash} #{edge.p2.hash}"
      manifest[edge.p1] ||= []
      manifest[edge.p2] ||= []
      manifest[edge.p1] << edge.p2
      manifest[edge.p2] << edge.p1
    end
  end
end