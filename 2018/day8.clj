(defn parse-input [input]
  (map #(Integer/parseInt %1) (re-seq #"\d+" input)))

(defn parse-partial-tree
  "Parses nums into a tree, return a vector of [tree rest-of-nums]"
  [nums]
  (let [num-child (first nums) num-meta (second nums)]
    (loop [nums (nthrest nums 2) n num-child children []]
      (if (zero? n)
        [{:children children :meta (take num-meta nums)} (nthrest nums num-meta)]
        (let [[child rest-of-nums] (parse-partial-tree nums)]
          (recur rest-of-nums (dec n) (conj children child)))))))

(defn parse-tree [nums]
  (first (parse-partial-tree nums)))

(defn sum-meta [tree]
  (apply + (concat (tree :meta) (map sum-meta (tree :children)))))

(defn tree-value [tree]
  (if (empty? (:children tree))
    (apply + (:meta tree))
    (let [children (tree :children)]
      (apply + (map tree-value (map #(nth children (dec %1) nil) (tree :meta)))))))

; (def nums (parse-input "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"))
(def nums (parse-input (slurp "day-8.txt")))
(def tree (parse-tree nums))
(println (sum-meta tree))
(println (tree-value tree))
