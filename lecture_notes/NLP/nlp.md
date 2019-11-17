#### Category 1. Architectures
- Traditional ML : "bag of words" => No order
- RNN : for sequences => Maintain order
- RNN 변형 : LSTM, GRU, ... 

#### Category 2. Embeddings
- Word2Vec, GloVe
- Gives words a location
- models relationships between words
- Word similarities 
- Analogies (King - man = Queen - woman)


Machine Translation -> Seq2Seq model 보통 사용 (input seq과 output seq. 길이 달라도 ok.)

#### Word Embeddings
- Why? weights와 biases (선형결합)를 학습하여 classification/regression 을 수행하기 위해서는
input이 string이 아닌 숫자 형태여야 함. => Word Embedding
- Sol#1 : One-hot encoding
- Feature Vector ? Data table에서 each row를 가리킴
- We want feature vectors to place things (points) in a <ins>meaningful position</ins> relative to each other.
- 3 Popular unsupervised methods : Word2Vec, GLoVe, FastText => Word vectors를 찾는 유명한 algorithms


Conventions : 
- V : # of rows = vocab. size = # of distinct words
- D = embedding dimension = feature vector size
- Word Embedding (Matrix) : V * D

When using Pre-trained weights :
- dataset이 slang과 같은, Pre-trained weights에 없는 단어를 포함하고 있다면?
- 해당 단어들을 initialize by random numbers (or 0)

#### Train Pre-trained embeddings? : 
- model.fit(X,Y) 는 model의 all parameters를 update 한다.
- pre-trained words에 포함되지 않았던 새로운 단어들은 randomly initialized 되었으므로, 학습이 필요하다.
- 그러나, 대부분의 library들은 updating only a subset of matrix를 허용하지 않음.. => Q. library 아니라면, 방법론이 나와 있을까?
(해결책 링크 : https://discuss.pytorch.org/t/update-part-of-word-embeddings-during-training/7617)
- library상에서는, 전체를 train하던가, 아니면 아예 하지 않던가 임.
- 실제 데이터에서는, fine-tuning을 하는 것이 나을 때도 있고, 아닐 때도 있음 => 둘 다 해봐야 함.
- Keras code :
  - To fine-fune : embedding_layer.trainable = True (default)
  - To keep it constant : embedding_layer.trainable = False
  
  cf. CNN에서의 Convolution은 원래 'Cross-Correlation' 이 정식 명칭이다.
  - 이미지에서 Edge detection 등의 목적의 filter로도 쓰이지만,
  - echo, reverb 등의 sound effect를 audio에 삽입할 때도 쓰인다.

