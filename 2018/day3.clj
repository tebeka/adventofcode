(use '[clojure.java.io :only (reader)])
(use '[clojure.set :only (union)])

(defn parse-line [line]
  (let [match (re-find #"(\d+) @ (\d+),(\d+): (\d+)x(\d+)" line)
        values (map #(Integer/parseInt %1) (rest match))
        [id x y width height] (into [] values)]
    {
     :id id
     :topX x
     :topY y
     :bottomX (+ x width)
     :bottomY (+ y height)
     }))

(defn overlap [c1 c2]
  (let [minX (max (c1 :topX) (c2 :topX))
        maxX (min (c1 :bottomX) (c2 :bottomX))
        minY (max (c1 :topY) (c2 :topY))
        maxY (min (c1 :bottomY) (c2 :bottomY))]
    (into #{} (for [x (range minX maxX) y (range minY maxY)] [x y]))))

(defn load-claims [path]
  (with-open [rdr (reader path)]
    (apply vector (map parse-line (line-seq rdr)))))

(defn claim-overlap [c claims]
  (let [others (remove #(= c %1) claims)]
    (reduce union (map #(overlap c %1) others))))

(defn part-1 [claims]
  (let [dups (reduce union (map #(claim-overlap %1 claims) claims))]
   (println (count dups))))


(defn part-2 [claims]
  (loop [i 0 claims claims]
    (let [c (claims i)]
      (if (zero? (count (claim-overlap c claims)))
        (println (c :id))
        (recur (inc i) claims)))))

(def claims (load-claims "day-3.txt"))
(part-1 claims)
(part-2 claims)


(comment
  (def c1 (parse-line "#1 @ 1,3: 4x4"))
  (def c2 (parse-line "#2 @ 3,1: 4x4"))
  (def c3 (parse-line "#3 @ 5,5: 2x2"))
  (def m [c1 c2 c3])
)
