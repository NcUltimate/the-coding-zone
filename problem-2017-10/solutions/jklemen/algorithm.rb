module Intersections
  class Algorithm
    def self.run!(raw1, raw2)
      self.new(raw1, raw2).run!
    end

    # polygons
    attr_accessor :raw1, :raw2, :augmented1, :augmented2

    # algorithm variables
    attr_accessor :visited, :children, :solution

    def initialize(raw1, raw2)
      self.raw1 = raw1
      self.raw2 = raw2

      self.visited = Set.new
      self.solution = []
      self.children = {}
    end

    def run!
      # 1. Augment polygons with intersections
      self.augmented1 = augmented_poly(raw1, raw2)
      self.augmented2 = augmented_poly(raw2, raw1)

      # 2. Starting at every point, try to find a path back to itself
      #     via points that intersect both polygons.
      start_points = []
      queue = initial_search_queue.to_a
      until queue.empty?
        nxt = queue.pop
        next if visited.member?(nxt)

        start_points << nxt
        visited.add(nxt)
        search(nxt)
      end

      # 3. Unfold the children of each point in 'start_points'
      #     to form new polygons and achieve the solution.
      start_points.each do |start|
        points = [start]
        curr = children[start]
        until curr.nil?
          points << curr
          curr = children[curr]
        end
        solution << points
      end

      solution
    end


    # Performs a depth-first search starting at a point by
    # grabbing all unvisited neighbors at the point, and
    # only iterating on each neighbor if
    #  1. That neighbor is unvisited
    #  2. The midpoint of the line between the point and
    #     the neighbor is also on both of the polygons
    # Before a neighbor is iterated on, it is set as the
    # child of the current point for purposes of discovering
    # connected components after the search is complete.
    def search(point)
      queue = unvisited_augmented_neighbors(point)
      until queue.empty?
        nxt = queue.pop
        next if visited.member?(nxt)

        line = Line.new(point, nxt)
        next unless augmented_intersection?(line.midpoint)

        visited.add(nxt)
        self.children[point] = nxt
        search(nxt)
      end
    end


    # Augments poly1 with points corresponding to
    # intersections with poly2. Note this does not change
    # the geometry, only separates existing edges into
    # smaller edges and more points
    def augmented_poly(poly1, poly2)
      points = poly1.edges.inject(Set.new) do |pts, edge|
        pts += intersections_for(poly2, edge)
      end
      Polygon.new(points.to_a)
    end


    # Finds all intersections of an edge in a polygon.
    # Sorts the result to return the points in order from
    # one side of the edge to the other instead of the order
    # they intersected in.
    def intersections_for(poly, edge)
      ints = poly.intersections(edge)
      ints.add(edge.p1)
      ints.add(edge.p2)
      ints.sort_by { |p| p.dist_from(edge.p1) }
    end


    # Starting points for the final phase of the algorithm.
    # We start with all of the points in both augmented
    # polygons and only choose the ones that intersect
    # both polygons.
    def initial_search_queue
      queue = (augmented1.points + augmented2.points)
      queue.select!(&method(:augmented_intersection?))
    end


    # Get neighbors for a point from both augmented polygons
    # as long as they have not been visited.
    def unvisited_augmented_neighbors(point)
      neighbors = Set.new
      neighbors += augmented1.neighbors[point] if augmented1.neighbors[point]
      neighbors += augmented2.neighbors[point] if augmented2.neighbors[point]
      neighbors.reject { |n| visited.member?(n) }
    end


    # Does this point fall within both augmented polygons?
    def augmented_intersection?(point)
      augmented1.intersects?(point) &&
        augmented2.intersects?(point)
    end
  end
end