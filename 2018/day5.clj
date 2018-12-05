(def cdiff (- (byte \a) (byte \A)))

(defn neg [c]
  (if (>= c (byte \a))
    (- c cdiff)
    (+ c cdiff)))


(defn reduce1 [polymer]
  (loop [polymer polymer out (vector)]
    (cond
      (empty? polymer) out
      (= (count polymer) 1) (conj out (first polymer))
      (= (first polymer) (neg (second polymer))) (recur (nthrest polymer 2) out)
      :else (recur (rest polymer) (conj out (first polymer))))))

(defn reduce-poly [polymer]
  (loop [polymer polymer]
    (let [out (reduce1 polymer)]
      (if (= (count polymer) (count out))
        polymer
        (recur out)))))

(defn part-1 [polymer]
  (println (count (reduce-poly polymer))))

(defn remove-char [c polymer]
  (filterv #(and (not= %1 c) (not= %1 (neg c))) polymer))

(defn part-2 [polymer]
  (loop [best -1 c (byte \a)]
    (if (> c (byte \z))
      (println best)
      (let [size (count (reduce-poly (remove-char c polymer)))]
        (if (or (neg? best) (< size best))
          (recur size (byte (inc c)))
          (recur best (byte (inc c))))))))

; pop remove newline
(def polymer (pop (mapv byte (slurp "day-5.txt"))))

(part-1 polymer)
(part-2 polymer)
