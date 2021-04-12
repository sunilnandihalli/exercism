(ns pov)

(defn of [pivot tree]
  (let [h (fn [[root & children :as tree]]
            (if (or (empty? tree) (= root pivot)) tree
                (loop [[cur-tree & siblings :as remaining-siblings] children prev-siblings []]
                  (if (empty? remaining-siblings) tree
                      (let [[proot :as ptree] (h cur-tree)]
                        (if (= proot pivot) (conj ptree (into [root] (into prev-siblings siblings)))
                            (recur siblings (conj prev-siblings cur-tree))))))))
        [root :as pivoted-tree] (h tree)]
    (if (= root pivot) pivoted-tree)))
#_(of :x [:x])
#_(of :x [:a [:x]])
#_(of :x [:a [:x] [:y]])
#_(of :x [:y])
#_(of :x [:a [:x [:b] [:c]] [:y]])
(defn path-from-to [] ;; <- arglist goes here
  ;; your code goes here
  )
