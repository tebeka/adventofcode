(use '[clojure.java.io :only (reader)])

(defn counts [s]
  (loop [m {} s s]
    (if (empty? s)
      (let [nums (set (vals m))]
        {2 (if (nums 2) 1 0) 3 (if (nums 3) 1 0)})
      (recur (update m (first s) (fnil inc 0)) (rest s)))))

(defn file-counts [path]
  (with-open [rdr (reader path)]
    (reduce (partial merge-with +) (map counts (line-seq rdr)))))

(defn part-1 []
  (let [counts (file-counts "day-2.txt")]
    (println (* (counts 2) (counts 3)))))

(part-1)

(defn diff-locs [s1 s2]
  "return vector of locations where s1 differ from s2"
  (let [ch (map vector (range (count s1)) s1 s2)]
    (apply vector (remove nil? (map #(if (not= (%1 1) (%1 2)) (%1 0)) ch)))))

(defn remove-char [s loc]
  (str (subs s 0 loc) (subs s (inc loc))))

(defn find-diff [s others]
  (loop [s s others others]
    (when (not (empty? others))
      (let [locs (diff-locs s (first others))]
        (if (= (count locs) 1)
          (remove-char s (locs 0))
          (recur s (rest others)))))))

(defn part-2 []
  (let [lines (with-open [rdr (reader "day-2.txt")] (doall (line-seq rdr)))]
    (println (first (remove nil? (map #(find-diff %1 lines) lines))))))

(part-2)
