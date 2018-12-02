(use '[clojure.java.io :only (reader)])

(defn part-1 []
  (with-open [rdr (reader "day-1.txt")]
    (println (reduce + (map read-string (line-seq rdr))))))

(part-1)

; part 2

(defn part-2 []
  (with-open [rdr (reader "day-1.txt")]
    (loop [nums (cycle (map read-string (line-seq rdr)))
           freq 0
           seen #{}]
      (let [next-freq (+ freq (first nums))]
        (if (seen next-freq)
          (println next-freq)
          (recur (rest nums) next-freq (conj seen next-freq)))))))

(part-2)
