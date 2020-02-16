#### Set up a problem -> what is Y?
#### Data collection -> collect X and Y
- (Traditionally) design a study, perform experiments, and collect data
- (Now) dig up a database and collect any related data 

#### Preprocessing
- Transform data into usable form 

#### Exploratory data analysis (EDA)
- See how data looks like 

#### Data analysis (prediction)
- Initial data size: n = 100M, p = 10k
- Separate training and test set (often by data collection time)
- Select features via univariate statistical test (t-test, cor-test)
- (optionally) Dimension reduction (PCA)
- Select learning methods (often based on intuition)
- Model selection via cross-validation (methods & parameters)
- Test performance over the test set

#### Validation with a totally new data set (often not existing data when the model is built)