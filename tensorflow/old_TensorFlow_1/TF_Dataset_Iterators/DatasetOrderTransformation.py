# Demonstrate the oredering of dataset transformation

import tensorflow as tf

dataset1 = tf.data.Dataset.from_tensor_slices(tf.range(10))
dataset1 = dataset1.batch(4)
dataset1 = dataset1.repeat(2)
dataset1 = dataset1.shuffle(4)


dataset2 = tf.data.Dataset.from_tensor_slices(tf.range(10))
dataset2 = dataset2.shuffle(4)
dataset2 = dataset2.repeat(2)
dataset2 = dataset2.batch(4)


# Code to try out data present in datasets
dataset = dataset2   # Change to required dataset
iterator = dataset.make_one_shot_iterator()
next_element = iterator.get_next()

with tf.Session() as sess:
  try:
    while True:
      print(sess.run(next_element))
  except:
    pass